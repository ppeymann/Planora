package repository

import (
	"github.com/jackc/pgconn"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/user/models"
	"gorm.io/gorm"
)

type userRepo struct {
	pg       *gorm.DB
	database string
	table    string
}

// Migrate implements models.UserRepository.
func (r *userRepo) Migrate() error {
	return r.pg.AutoMigrate(&models.UserEntity{})
}

// Model implements models.UserRepository.
func (r *userRepo) Model() *gorm.DB {
	return r.pg.Model(&models.UserEntity{})
}

// Name implements models.UserRepository.
func (r *userRepo) Name() string {
	return r.table
}

// Create implements models.UserRepository.
func (r *userRepo) Create(in *userpb.SignUpRequest) (*models.UserEntity, error) {
	user := &models.UserEntity{
		Model:     gorm.Model{},
		Username:  in.Username,
		Password:  in.Password,
		Email:     in.Email,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	// create Account
	err := r.pg.Transaction(func(tx *gorm.DB) error {
		if err := r.Model().Create(user).Error; err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok {
				if pgErr.Code == "23505" {
					return models.ErrAccountExist
				}
			}
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRepo(db *gorm.DB, database string) models.UserRepository {
	return &userRepo{
		pg:       db,
		database: database,
		table:    "user_entities",
	}
}

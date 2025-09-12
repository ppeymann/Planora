package repository

import (
	"time"

	"github.com/jackc/pgconn"

	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/user/models"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type userRepo struct {
	pg       *gorm.DB
	database string
	table    string
}

// GetRoomUsers implements models.UserRepository.
func (r *userRepo) GetRoomUsers(ids []uint64) ([]models.UserEntity, error) {
	var users []models.UserEntity
	if len(ids) == 0 {
		return users, nil
	}

	uids := make([]uint, len(ids))
	for i, id := range ids {
		uids[i] = uint(id)
	}

	if err := r.Model().Where("id IN ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// FindByID implements models.UserRepository.
func (r *userRepo) FindByID(id uint) (*models.UserEntity, error) {
	user := &models.UserEntity{}

	err := r.Model().Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Find implements models.UserRepository.
func (r *userRepo) Find(username string) (*models.UserEntity, error) {
	user := &models.UserEntity{}

	if err := r.Model().Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Update implements models.UserRepository.
func (r *userRepo) Update(user *models.UserEntity) error {
	return r.pg.Save(user).Error
}

// Migrate implements models.UserRepository.
func (r *userRepo) Migrate() error {
	err := r.pg.AutoMigrate(&models.RefreshTokenEntity{})
	if err != nil {
		return err
	}

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

	refresh := models.RefreshTokenEntity{
		TokenId:   ksuid.New().String(),
		IssuedAt:  time.Now().UTC().Unix(),
		ExpiredAt: time.Now().Add(time.Duration(1036800 * time.Minute)).UTC().Unix(),
	}

	user.Tokens = append(user.Tokens, refresh)

	err = r.Update(user)
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

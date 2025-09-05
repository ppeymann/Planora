package repository

import (
	todopb "github.com/ppeymann/Planora.git/proto/todo"
	"github.com/ppeymann/Planora/todo/models"
	"gorm.io/gorm"
)

type todoRepo struct {
	pg       *gorm.DB
	database string
	table    string
}

// Create implements models.TodoRepository.
func (r *todoRepo) Create(in *todopb.AddTodoRequest) (*models.TodoEntity, error) {
	todo := &models.TodoEntity{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
		Status:      models.StatusDo,
		UserID:      uint(in.GetUserId()),
	}

	if err := r.Model().Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// Migrate implements models.TodoRepository.
func (r *todoRepo) Migrate() error {
	return r.pg.AutoMigrate(&models.TodoEntity{})
}

// Model implements models.TodoRepository.
func (r *todoRepo) Model() *gorm.DB {
	return r.pg.Model(&models.TodoEntity{})
}

// Name implements models.TodoRepository.
func (r *todoRepo) Name() string {
	return r.table
}

func NewTodoRepo(db *gorm.DB, database string) models.TodoRepository {
	return &todoRepo{
		pg:       db,
		database: database,
		table:    "todo_entities",
	}
}

package models

import (
	"github.com/ppeymann/Planora.git/pkg/common"
	todopb "github.com/ppeymann/Planora.git/proto/todo"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type (
	EventType string

	StatusType string

	TodoRepository interface {
		Create(in *todopb.AddTodoRequest) (*TodoEntity, error)

		common.BaseRepository
	}

	TodoEntity struct {
		gorm.Model

		// Title
		Title string `json:"title" gorm:"column:title;index;not null"`

		// Description
		Description string `json:"description" gorm:"column:description"`

		// Status
		Status StatusType `json:"status" gorm:"column:status;default:'DO';check:status IN ('DO','IN_PROGRESS','DONE','ARCHIVED','BLOCKED')"`

		// UserID
		UserID uint `json:"user_id" gorm:"column:user_id;index;not null"`
	}
)

const (
	StatusDo         StatusType = "DO"
	StatusInProgress StatusType = "IN_PROGRESS"
	StatusDone       StatusType = "DONE"
	StatusArchived   StatusType = "ARCHIVED"
	StatusBlocked    StatusType = "BLOCKED"
)

const (
	SubjectAddTodo EventType = "todo.ADD"
)

func ToBaseModel(t *TodoEntity) *todopb.BaseModel {
	return &todopb.BaseModel{
		Id:         uint64(t.ID),
		CreatedAt:  timestamppb.New(t.CreatedAt),
		UpdatedeAt: timestamppb.New(t.UpdatedAt),
	}
}

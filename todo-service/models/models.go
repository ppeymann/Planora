package models

import (
	"github.com/ppeymann/Planora.git/pkg/common"
	"gorm.io/gorm"
)

type (
	EventType string

	StatusType string

	TodoRepository interface {
		Create() (*TodoEntity, error)

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
	}
)

const (
	StatusDo         StatusType = "DO"
	StatusInProgress StatusType = "IN_PROGRESS"
	StatusDone       StatusType = "DONE"
	StatusArchived   StatusType = "ARCHIVED"
	StatusBlocked    StatusType = "BLOCKED"
)

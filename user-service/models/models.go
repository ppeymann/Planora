package models

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/ppeymann/Planora.git/pkg/common"

	userpb "github.com/ppeymann/Planora.git/proto/user"
)

var (
	ErrAccountExist error = errors.New("account with specified params already exists")
)

type (
	EventType string

	UserRepository interface {
		// Create new user
		Create(in *userpb.SignUpRequest) (*UserEntity, error)

		// Update user
		Update(user *UserEntity) error

		common.BaseRepository
	}

	UserEntity struct {
		gorm.Model

		Username  string               `json:"username" gorm:"uniqueIndex;not null"`
		Email     string               `json:"email" gorm:"uniqueIndex;not null"`
		Password  string               `json:"-" gorm:"not null"`
		FirstName string               `json:"first_name" gorm:"column:first_name"`
		LastName  string               `json:"last_name" gorm:"column:last_name"`
		Tokens    []RefreshTokenEntity `json:"-" gorm:"foreignKey:AccountID;references:ID"`
	}

	RefreshTokenEntity struct {
		gorm.Model
		AccountID uint
		TokenId   string `json:"token_id" gorm:"column:token_id;index"`
		UserAgent string `json:"user_agent" gorm:"column:user_agent"`
		IssuedAt  int64  `json:"issued_at" bson:"issued_at" gorm:"column:issued_at"`
		ExpiredAt int64  `json:"expired_at" bson:"expired_at" gorm:"column:expired_at"`
	}

	TokenBundlerOutput struct {
		// Token is string that hashed by paseto
		Token string `json:"token"`

		// Refresh is string that for refresh old token
		Refresh string `json:"refresh"`

		// Expire is time for expire token
		Expire time.Time `json:"expire"`
	}
)

const (
	Signup EventType = "user.SIGNUP"
)

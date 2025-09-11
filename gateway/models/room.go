package models

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
)

type (
	// RoomService represents method signatures for api room endpoint.
	// so any object that stratifying this interface can be used as room service for api endpoint.
	RoomService interface {
		// Create is service for create room
		Create(ctx *gin.Context, in *RoomInput) *common.BaseResult
	}

	// RoomHandler represents method signatures for room handlers.
	// so any object that stratifying this interface can be used as room handlers.
	RoomHandler interface {
		// Create is handler for create room
		Create(ctx *gin.Context)
	}

	RoomInput struct {
		Name      string `json:"name"`
		CreatorID uint64 `json:"creator_id"`
	}
)

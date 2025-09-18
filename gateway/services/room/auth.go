package room

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/auth"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora.git/pkg/utils"
	"github.com/ppeymann/Planora/gateway/models"
)

type authService struct {
	next models.RoomService
}

// AddUser implements models.RoomService.
func (a *authService) AddUser(ctx *gin.Context, in *models.AddUserInput) *common.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{common.ErrUnAuthorization.Error()},
			Status: http.StatusUnauthorized,
		}
	}

	in.CreatorID = uint64(claims.Subject)
	return a.next.AddUser(ctx, in)
}

// GetRoom implements models.RoomService.
func (a *authService) GetRoom(ctx *gin.Context, roomID uint64) *common.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{common.ErrUnAuthorization.Error()},
			Status: http.StatusUnauthorized,
		}
	}

	return a.next.GetRoom(ctx, roomID)
}

// Create implements models.RoomService.
func (a *authService) Create(ctx *gin.Context, in *models.RoomInput) *common.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{common.ErrUnAuthorization.Error()},
			Status: http.StatusUnauthorized,
		}
	}

	in.CreatorID = uint64(claims.Subject)
	return a.next.Create(ctx, in)
}

func NewAuthService(srv models.RoomService) models.RoomService {
	return &authService{
		next: srv,
	}
}

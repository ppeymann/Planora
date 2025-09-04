package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/auth"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora.git/pkg/utils"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/gateway/models"
)

type authService struct {
	next models.UserService
}

// Account implements models.UserService.
func (a *authService) Account(ctx *gin.Context) *common.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{common.ErrUnAuthorization.Error()},
			Status: http.StatusUnauthorized,
		}
	}

	return a.next.Account(ctx)
}

// Login implements models.UserService.
func (a *authService) Login(ctx *gin.Context, in *userpb.LoginRequest) *common.BaseResult {
	return a.next.Login(ctx, in)
}

// SignUp implements models.UserService.
func (a *authService) SignUp(ctx *gin.Context, in *userpb.SignUpRequest) *common.BaseResult {
	return a.next.SignUp(ctx, in)
}

func NewAuthService(service models.UserService) models.UserService {
	return &authService{
		next: service,
	}
}

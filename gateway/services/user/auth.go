package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/gateway/models"
)

type authService struct {
	next models.UserService
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

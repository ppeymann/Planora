package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/auth"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora.git/pkg/utils"
	"github.com/ppeymann/Planora/gateway/models"
)

type authService struct {
	next models.TodoService
}

// UpdateTodo implements models.TodoService.
func (a *authService) UpdateTodo(ctx *gin.Context, in *models.TodoInput, todoID uint64) *common.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{common.ErrUnAuthorization.Error()},
			Status: http.StatusUnauthorized,
		}
	}

	in.UserID = claims.Subject
	return a.next.UpdateTodo(ctx, in, todoID)
}

// AddTodo implements models.TodoService.
func (a *authService) AddTodo(ctx *gin.Context, in *models.TodoInput) *common.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{common.ErrUnAuthorization.Error()},
			Status: http.StatusUnauthorized,
		}
	}

	in.UserID = claims.Subject
	return a.next.AddTodo(ctx, in)
}

func NewAuthService(s models.TodoService) models.TodoService {
	return &authService{
		next: s,
	}
}

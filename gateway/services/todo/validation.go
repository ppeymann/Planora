package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/gateway/models"
	validations "github.com/ppeymann/Planora/gateway/validation"
)

type validationService struct {
	schemas map[string][]byte
	next    models.TodoService
}

// DeleteTodo implements models.TodoService.
func (v *validationService) DeleteTodo(ctx *gin.Context, id uint64) *common.BaseResult {
	return v.next.DeleteTodo(ctx, id)
}

// ChangeStatus implements models.TodoService.
func (v *validationService) ChangeStatus(ctx *gin.Context, status models.StatusType, id uint64) *common.BaseResult {
	return v.next.ChangeStatus(ctx, status, id)
}

// GetAllTodos implements models.TodoService.
func (v *validationService) GetAllTodos(ctx *gin.Context) *common.BaseResult {
	return v.next.GetAllTodos(ctx)
}

// UpdateTodo implements models.TodoService.
func (v *validationService) UpdateTodo(ctx *gin.Context, in *models.TodoInput, todoID uint64) *common.BaseResult {
	err := validations.Validate(in, v.schemas)
	if err != nil {
		return err
	}

	return v.next.UpdateTodo(ctx, in, todoID)
}

// AddTodo implements models.TodoService.
func (v *validationService) AddTodo(ctx *gin.Context, in *models.TodoInput) *common.BaseResult {
	err := validations.Validate(in, v.schemas)
	if err != nil {
		return err
	}

	return v.next.AddTodo(ctx, in)
}

func NewValidationService(path string, s models.TodoService) (models.TodoService, error) {
	schemas := make(map[string][]byte)

	err := validations.LoadSchema(path, schemas)
	if err != nil {
		return nil, err
	}

	return &validationService{
		schemas: schemas,
		next:    s,
	}, nil
}

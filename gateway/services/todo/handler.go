package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/gateway/models"
	"github.com/ppeymann/Planora/gateway/server"
)

type handler struct {
	next models.TodoService
}

// ChangeStatus implements models.TodoHandler.
func (h *handler) ChangeStatus(ctx *gin.Context) {
	status, err := server.GetStringPath("status", ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResult{
			Errors: []string{err.Error()},
		})

		return
	}

	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResult{
			Errors: []string{err.Error()},
		})

		return
	}

	result := h.next.ChangeStatus(ctx, models.StatusType(status), id)
	ctx.JSON(result.Status, result)
}

// GetAllTodos implements models.TodoHandler.
func (h *handler) GetAllTodos(ctx *gin.Context) {
	result := h.next.GetAllTodos(ctx)
	ctx.JSON(result.Status, result)
}

// UpdateTodo implements models.TodoHandler.
func (h *handler) UpdateTodo(ctx *gin.Context) {
	in := &models.TodoInput{}

	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResult{
			Errors: []string{models.ErrProvideRequiredJsonBody.Error()},
		})

		return
	}

	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResult{
			Errors: []string{err.Error()},
		})

		return
	}

	result := h.next.UpdateTodo(ctx, in, id)
	ctx.JSON(result.Status, result)
}

// AddTodo implements models.TodoHandler.
func (h *handler) AddTodo(ctx *gin.Context) {
	in := &models.TodoInput{}

	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResult{
			Errors: []string{models.ErrProvideRequiredJsonBody.Error()},
		})

		return
	}

	result := h.next.AddTodo(ctx, in)
	ctx.JSON(result.Status, result)
}

func NewHandler(srv models.TodoService, s *server.Server) models.TodoHandler {
	handler := &handler{
		next: srv,
	}

	group := s.Router.Group("/api/v1/todo")

	group.Use(s.Authenticate())
	{
		group.POST("/", handler.AddTodo)
		group.PATCH("/:id", handler.UpdateTodo)
		group.GET("/", handler.GetAllTodos)
	}

	return handler
}

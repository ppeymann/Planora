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
	}

	return handler
}

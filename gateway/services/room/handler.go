package room

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/gateway/models"
	"github.com/ppeymann/Planora/gateway/server"
)

type handler struct {
	next models.RoomService
}

// Create implements models.RoomHandler.
func (h *handler) Create(ctx *gin.Context) {
	in := &models.RoomInput{}

	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResult{
			Errors: []string{models.ErrProvideRequiredJsonBody.Error()},
		})

		return
	}

	result := h.next.Create(ctx, in)
	ctx.JSON(result.Status, result)
}

func NewHandler(srv models.RoomService, s *server.Server) models.RoomHandler {
	handler := &handler{
		next: srv,
	}

	group := s.Router.Group("/api/v1/room")
	group.Use(s.Authenticate())
	{
		group.POST("/", handler.Create)
	}

	return handler
}

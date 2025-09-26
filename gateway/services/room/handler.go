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

// AddUser implements models.RoomHandler.
func (h *handler) AddUser(ctx *gin.Context) {
	in := &models.AddUserInput{}
	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResult{
			Errors: []string{models.ErrProvideRequiredJsonBody.Error()},
		})

		return
	}

	result := h.next.AddUser(ctx, in)
	ctx.JSON(result.Status, result)
}

// GetRoom implements models.RoomHandler.
func (h *handler) GetRoom(ctx *gin.Context) {
	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResult{
			Errors: []string{err.Error()},
		})

		return
	}

	result := h.next.GetRoom(ctx, id)
	ctx.JSON(result.Status, result)
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
		group.GET("/:id", handler.GetRoom)
		group.POST("/add_user", handler.AddUser)
	}

	return handler
}

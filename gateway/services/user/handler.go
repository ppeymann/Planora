package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/gateway/models"
	"github.com/ppeymann/Planora/gateway/server"
)

type handler struct {
	next models.UserService
}

// SignUp implements models.UserHandler.
func (h *handler) SignUp(ctx *gin.Context) {
	in := &userpb.SignUpRequest{}

	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResult{
			Errors: []string{models.ErrProvideRequiredJsonBody.Error()},
		})

		return
	}

	result := h.next.SignUp(ctx, in)
	ctx.JSON(result.Status, result)
}

func NewHandler(service models.UserService, s *server.Server) models.UserHandler {
	handler := &handler{
		next: service,
	}

	group := s.Router.Group("/api/v1/user")
	{
		group.POST("/signup", handler.SignUp)

	}

	return handler
}

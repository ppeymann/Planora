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

// Login implements models.UserHandler.
//
// @BasePath			/api/v1/user
// @Summary				login user
// @Description			login user if signed up
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				input	body	userpb.LoginRequest		 true	"login input"
// @Success				200		{object}	common.BaseResult	"always return status 200 but body contains errors"
// @Router				/login	[post]
func (h *handler) Login(ctx *gin.Context) {
	in := &userpb.LoginRequest{}

	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, &common.BaseResult{
			Errors: []string{models.ErrProvideRequiredJsonBody.Error()},
		})

		return
	}

	result := h.next.Login(ctx, in)
	ctx.JSON(result.Status, result)
}

// SignUp implements models.UserHandler.
//
// @BasePath			/api/v1/user
// @Summary				sign up user
// @Description			sign up user if never logged in
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				input	body	userpb.SignUpRequest 	true	"sign up input"
// @Success				200		{object}	common.BaseResult	"always return status 200 but body contains errors"
// @Router				/signup	[post]
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
		group.POST("/login", handler.Login)

	}

	return handler
}

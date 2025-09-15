package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/gateway/models"
	validations "github.com/ppeymann/Planora/gateway/validation"
)

type validationService struct {
	next   models.UserService
	schema map[string][]byte
}

// Account implements models.UserService.
func (v *validationService) Account(ctx *gin.Context) *common.BaseResult {
	return v.next.Account(ctx)
}

// Login implements models.UserService.
func (v *validationService) Login(ctx *gin.Context, in *models.LoginInput) *common.BaseResult {
	err := validations.Validate(in, v.schema)
	if err != nil {
		return err
	}

	return v.next.Login(ctx, in)
}

// SignUp implements models.UserService.
func (v *validationService) SignUp(ctx *gin.Context, in *models.SignUpInput) *common.BaseResult {
	err := validations.Validate(in, v.schema)
	if err != nil {
		return err
	}

	return v.next.SignUp(ctx, in)
}

func NewValidationService(srv models.UserService, path string) (models.UserService, error) {
	schema := make(map[string][]byte)

	// Load the schema from the specified path
	err := validations.LoadSchema(path, schema)
	if err != nil {
		return nil, err
	}

	return &validationService{
		next:   srv,
		schema: schema,
	}, nil
}

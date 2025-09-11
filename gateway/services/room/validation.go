package room

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/gateway/models"
	validations "github.com/ppeymann/Planora/gateway/validation"
)

type validationService struct {
	schemas map[string][]byte
	next    models.RoomService
}

// Create implements models.RoomService.
func (v *validationService) Create(ctx *gin.Context, in *models.RoomInput) *common.BaseResult {
	err := validations.Validate(in, v.schemas)
	if err != nil {
		return err
	}

	return v.next.Create(ctx, in)
}

func NewValidationService(path string, srv models.RoomService) (models.RoomService, error) {
	schemas := make(map[string][]byte)

	err := validations.LoadSchema(path, schemas)
	if err != nil {
		return nil, err
	}

	return &validationService{
		schemas: schemas,
		next:    srv,
	}, nil
}

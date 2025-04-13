package jwt_secret

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmhtoo/go-logger-api/config"
	"github.com/mmhtoo/go-logger-api/helpers"
)

type JwtSecretHandler struct {
	jwtSecretService *JwtSecretService
}

func NewJwtSecretHandler(database *config.Database) *JwtSecretHandler {
	return &JwtSecretHandler{
		jwtSecretService: NewJwtSecretService(
			NewJwtSecretRepository(database),
		),
	}
}

func (h *JwtSecretHandler) HandleGetAllJwtSecretsByProjectId(ctx *gin.Context){
	projectId, isOk := ctx.Params.Get("id")
	if !isOk {
		ctx.JSON(
			http.StatusBadRequest,
			helpers.NewAPIErrorResponse("Invalid project id!", "Validation failed!"),
		)
		return
	}
	jwtSecrets, err := h.jwtSecretService.GetAllJwtSecretsByProjectId(
		projectId,
		ctx.Request.Context(),
	)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			helpers.NewAPIErrorResponse(
				err.Error(),
				"Failed to get jwt secrets!",
			),
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		helpers.NewAPIDataResponse(
			MapJwtSecretEntitesToResDto(jwtSecrets), "Successfully retrieved!",
		),
	)
}

func (h *JwtSecretHandler) HandleGetDetailById(ctx *gin.Context){
	secretId, isOk := ctx.Params.Get("secretId")
	if !isOk {
		ctx.JSON(
			http.StatusBadRequest,
			helpers.NewAPIErrorResponse("Invalid secret id!", "Validation failed!"),
		)
		return
	}
	jwtSecret, err := h.jwtSecretService.GetDetailById(secretId, ctx.Request.Context())
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			helpers.NewAPIErrorResponse(err.Error(), "Failed to get jwt secret!"),
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		helpers.NewAPIDataResponse(jwtSecret.ToDetailResponseDto(), "Successfully retrieved!"),
	)
}
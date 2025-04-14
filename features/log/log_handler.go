package log

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmhtoo/go-logger-api/config"
	"github.com/mmhtoo/go-logger-api/helpers"
)

type LogHandler struct {
	logService *LogService
}

func NewLogHandler(database *config.Database) *LogHandler {
	return &LogHandler{
		logService: NewLogService(NewLogRepository(database)),
	}
}

func (h *LogHandler) HandleSaveLog(c *gin.Context){
	reqContext := c.Request.Context()
	reqBody := c.MustGet("payload").(*SaveLogReqDto)
	err := h.logService.Save(
		reqBody.ToSaveLogInput(),
		reqContext,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			helpers.NewAPIErrorResponse(
				err.Error(),
				"Failed to save log!",
			),
		)
		return
	}
	c.JSON(http.StatusCreated, helpers.NewAPIBaseResponse("Success!"))	
}

func (h* LogHandler) HandleGetLogsWithFilter(c *gin.Context){
	ctx := c.Request.Context()
	queryPayload, exists := c.Get("query")
	if !exists {
		c.JSON(
			http.StatusBadRequest,
			helpers.NewAPIErrorResponse(
				errors.New("Invalid query payload!"),
				"Validation failed!",
			),
		)
		return
	}
	logs, err := h.logService.GetLogsWithFilter(
		(queryPayload.(*GetLogsWithFilterReqDto)).ToSelectByProjectIdWithFilterInput(),
		ctx,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			helpers.NewAPIErrorResponse(
				err.Error(),
				"Failed to process!",
			),
		)
		return
	}
	c.JSON(
		http.StatusOK,
		helpers.NewAPIDataResponse(
			MapLogEntitiesToResDto(logs),
			"Successfully retrieved!",
		),
	)
}
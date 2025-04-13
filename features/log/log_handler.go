package log

import (
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
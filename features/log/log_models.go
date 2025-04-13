package log

import (
	"time"

	"github.com/google/uuid"
)

type LogEntity struct {
	id string `db:"id"`
	logType string `db:"log_type"`
	loggedAt time.Time `db:"logged_at"`
	loggedBy string `db:"logged_by"`
	pathName string `db:"path_name"`
	projectId string `db:"project_id"`
	payload string `db:"payload"`
}

type SaveLogInput struct {
	Id string 
	LogType string
	LoggedAt time.Time
	LoggedBy string
	PathName string
	ProjectId string
	Payload string	
}

type SaveLogReqDto struct {
	LogType string `json:"logType" binding:"required,min=3,max=100,oneofci=info debug error"`
	PathName string `json:"pathName" binding:"required,min=3,max=255"`
	ProjectId string `json:"projectId" binding:"required,min=3,max=255"`
	Payload string `json:"payload" binding:"required,json"`
	SecretKeyId string `json:"secretKeyId" binding:"required"`
}

func (dto *SaveLogReqDto) ToSaveLogInput() *SaveLogInput {
	return &SaveLogInput{
		Id: uuid.NewString(),
		LogType: dto.LogType,
		LoggedAt: time.Now(),
		LoggedBy: dto.SecretKeyId,
		PathName: dto.PathName,
		ProjectId: dto.ProjectId,
		Payload: dto.Payload,
	}
}

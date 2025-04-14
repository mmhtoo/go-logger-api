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

type SelectByProjectIdWithFilterInput struct {
	projectId string
	logType string
	fromTime time.Time
	toTime time.Time
	pathName string 
	keyword string
	page int16
	pageSize int16
}

type GetLogsWithFilterReqDto struct {
	ProjectId string `form:"projectId" binding:"required"`
	LogType string `form:"logType" binding:"omitempty,oneofci=info debug error"`
	FromTime time.Time `form:"fromTime" time_format:"2006-01-02T15:04:05" binding:"required"`
	ToTime time.Time `form:"toTime" time_format:"2006-01-02T15:04:05" binding:"omitempty"`
	PathName string `form:"pathName" binding:"omitempty"`
	Keyword string `form:"keyword" binding:"omitempty"`
	Page int16 `form:"page" binding:"required,min=1"`
	PageSize int16 `form:"pageSize" binding:"required,min=1"`
}

type LogResponseDto struct {
	Id string `json:"id"`
	LogType string `json:"logType"`
	LoggedAt time.Time `json:"loggedAt"`
	LoggedBy string `json:"loggedBy"`
	PathName string `json:"pathName"`
	ProjectId string `json:"projectId"`	
	Payload string `json:"payload"`
}

func (entity *LogEntity) ToResponseDto() *LogResponseDto {
	return &LogResponseDto{
		Id: entity.id,
		LogType: entity.logType,
		LoggedAt: entity.loggedAt,
		LoggedBy: entity.loggedBy,
		PathName: entity.pathName,
		ProjectId: entity.projectId,
		Payload: entity.payload,
	}
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

func (dto *GetLogsWithFilterReqDto) ToSelectByProjectIdWithFilterInput() *SelectByProjectIdWithFilterInput {
	return &SelectByProjectIdWithFilterInput{
		projectId: dto.ProjectId,
		logType: dto.LogType,
		fromTime: dto.FromTime,
		toTime: dto.ToTime,
		pathName: dto.PathName,
		keyword: dto.Keyword,
		page: dto.Page,
		pageSize: dto.PageSize,
	}
}

func MapLogEntitiesToResDto(entities *[]LogEntity) *[]LogResponseDto{
	var dtos []LogResponseDto
	for _, entity := range *entities {
		dtos = append(dtos, *entity.ToResponseDto())
	}
	return &dtos
}
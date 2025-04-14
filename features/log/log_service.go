package log

import "context"

type LogService struct {
	logRepository *LogReposistory
}

func NewLogService(logRepository *LogReposistory) *LogService {
	return &LogService{
		logRepository: logRepository,
	}
}

func (service *LogService) Save(input *SaveLogInput, ctx context.Context) error {
	return service.logRepository.Save(input, ctx)
}

func (service *LogService) GetLogsWithFilter(
	input *SelectByProjectIdWithFilterInput, 
	ctx context.Context,
) (*[]LogEntity, error) {
	logs, err := service.logRepository.SelectByProjectIdWithFilter(input, ctx)
	return logs, err
}
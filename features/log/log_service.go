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
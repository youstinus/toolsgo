package services

import (
	"github.com/sirupsen/logrus"
	"github.com/youstinus/toolsgo/pkg/config"
	"github.com/youstinus/toolsgo/pkg/services/examples_service"
)

// If describes all services
type If interface {
	GetExamplesService() examples_service.If
}

// Services holds services
type Services struct {
	ExamplesService examples_service.If
}

// Init creates services instance
func Init(cfg *config.Services, log *logrus.Logger) If {
	return &Services{
		ExamplesService: examples_service.Init(&cfg.ExamplesService, log),
	}
}

// GetExamplesService returns service instance
func (s *Services) GetExamplesService() examples_service.If {
	return s.ExamplesService
}

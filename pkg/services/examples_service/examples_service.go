package examples_service

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/youstinus/toolsgo/pkg/config"
	"github.com/youstinus/toolsgo/pkg/types"
)

//go:generate mockgen -source=examples_service.go -destination=mock_examples_service/mock_examples_service.go

const (
	// all constant strings.
	titleExampleNotFound = "example was not found"
)

var (
	// package related variables.
	errExampleNotFound = errors.New(titleExampleNotFound)
)

// If describes service functionality.
type If interface {
	// GetExample gets example object using given ID.
	GetExample(exampleID int64) (*types.Example, error)
}

// ExamplesService holds private fields.
type ExamplesService struct {
	log     *logrus.Logger
	example string
}

// Init creates service using values.
func Init(cfg *config.ExamplesService, log *logrus.Logger) If {
	return &ExamplesService{
		log:     log,
		example: cfg.Example,
	}
}

// GetExample contains all business logic.
func (s *ExamplesService) GetExample(exampleID int64) (*types.Example, error) {
	example := &types.Example{}
	return example, nil
}

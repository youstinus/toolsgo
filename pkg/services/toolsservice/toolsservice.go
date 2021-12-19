package toolsservice

import (
	"encoding/base64"

	"github.com/sirupsen/logrus"
	"github.com/youstinus/toolsgo/pkg/config"
)

// If describes service functionality.
type If interface {
	// EncodeBase64 returns encoded string.
	EncodeBase64(input string) string
	// DecodeBase64 decodes encoded string.
	// returns error if given string is not base64 and canont be decoded.
	DecodeBase64(input string) (string, error)
}

// Service holds private fields.
type Service struct {
	log *logrus.Logger
}

// Init creates service using values.
func Init(cfg *config.ToolsService, log *logrus.Logger) If {
	return &Service{
		log: log,
	}
}

func (s *Service) EncodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func (s *Service) DecodeBase64(input string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

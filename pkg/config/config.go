package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Logger configurations for logging
type Logger struct {
	LogLevel     string `yaml:"LogLevel"`     // LogLevels can be found in github.com/sirupsen/logrus/logrus.go:25 (ParseLevel method)
	LogFormatter string `yaml:"LogFormatter"` // LogFormatters can be "json" or "text"
}

// HTTP holds REST server API info
type HTTP struct {
	Enabled       bool   `yaml:"Enabled"`
	Address       string `yaml:"Address"`
	Port          int    `yaml:"Port"`
	Key           string `yaml:"Key"`
	EnableCors    bool   `yaml:"EnableCors"`
	EnableMetrics bool   `yaml:"EnableMetrics"`
}

// ExamplesService configurations
type ExamplesService struct {
	Example string `yaml:"Example"`
}

// Services holds info for all services
type Services struct {
	ExamplesService `yaml:"ExamplesService"`
}

// Config contains whole app configurations
type Config struct {
	Logger   `yaml:"Logger"`
	HTTP     `yaml:"HTTP"`
	Services `yaml:"Services"`
}

// ReadYAMLfile
func ReadYAMLfile(fileName string) *Config {
	yamlData, err := ioutil.ReadFile(fileName)
	if nil != err {
		log.Fatalf("error while Reading file:%s", err.Error())
	}
	cfg := Config{}
	err = yaml.Unmarshal(yamlData, &cfg)
	if nil != err {
		log.Fatalf("error while Unmarshal file:%s", err.Error())
	}
	return &cfg
}

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/youstinus/toolsgo/api/server"
	"github.com/youstinus/toolsgo/api/server/router"
	"github.com/youstinus/toolsgo/pkg/config"
)

// MakeHTTPServer creates server.Options and passes to server.New(). Creates Server with private parameters.
func MakeHTTPServer(cfg *config.HTTP, log *logrus.Logger, controllers router.If) *server.Server {
	opts := &server.Options{
		Enabled:       cfg.Enabled,
		Log:           log,
		HTTPPort:      cfg.Port,
		EnableCors:    cfg.EnableCors,
		EnableMetrics: cfg.EnableMetrics,
		AuthKey:       cfg.Key,
		Controllers:   controllers,
	}

	return server.New(opts)
}

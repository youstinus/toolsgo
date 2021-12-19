package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/youstinus/toolsgo/api/server/router/examples"
	"github.com/youstinus/toolsgo/pkg/services"
)

// If used to access all handlers.
type If interface {
	InitExamplesRoutes(r *chi.Mux)
}

// Controllers holds all handlers.
type Controllers struct {
	Controller examples.If
}

// InitControllers creates controllers using services.
func InitControllers(s services.If) If {
	return &Controllers{
		Controller: examples.Init(s.GetExamplesService()),
	}
}

// InitExamplesRoutes initializes routes with controllers.
func (c *Controllers) InitExamplesRoutes(r *chi.Mux) {
	c.Controller.InitRoutes(r)
}

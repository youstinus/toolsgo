package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/youstinus/toolsgo/api/server/router/examples"
	"github.com/youstinus/toolsgo/api/server/router/tools"
	"github.com/youstinus/toolsgo/pkg/services"
)

// If used to access all handlers.
type If interface {
	// InitRoutes initializes routes with controllers.
	InitRoutes(r *chi.Mux)
}

// Controllers holds all handlers.
type Controllers struct {
	Controller examples.If
	Tools      tools.If
}

// InitControllers creates controllers using services.
func InitControllers(s services.If) If {
	return &Controllers{
		Controller: examples.Init(s.GetExamplesService()),
		Tools:      tools.Init(s.GetToolsService()),
	}
}

func (c *Controllers) InitRoutes(r *chi.Mux) {
	c.Controller.InitRoutes(r)
	c.Tools.InitRoutes(r)
}

package tools

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/youstinus/toolsgo/pkg/services/toolsservice"
)

// strings used in package
const (
	constToolsURL = "/tools"
)

// If describes handler methods for Tools endpoint
type If interface {
	// InitRoutes initializes Tools routes.
	InitRoutes(r *chi.Mux)

	// encodeBase64 encodes data given in body.
	encodeBase64(w http.ResponseWriter, r *http.Request)
	// decodeBase64 decodes data given in body.
	decodeBase64(w http.ResponseWriter, r *http.Request)
}

// Controller holds tools service
type Controller struct {
	toolsService toolsservice.If
}

// Init creates Controller using service
func Init(toolsService toolsservice.If) If {
	return &Controller{
		toolsService: toolsService,
	}
}

func (c *Controller) InitRoutes(r *chi.Mux) {
	r.Post(constToolsURL+"/encodebase64", c.encodeBase64)
	r.Post(constToolsURL+"/decodebase64", c.decodeBase64)
}

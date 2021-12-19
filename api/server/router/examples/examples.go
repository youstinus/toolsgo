package examples

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/youstinus/toolsgo/pkg/services/examples_service"
)

// strings used in package
const (
	constExamplesURL    = "/examples"
	constExampleID      = "exampleID"
	constExampleIDParam = "/{exampleID}"
	constExampleURL     = constExamplesURL + constExampleIDParam
)

// If describes handler methods for Examples endpoint
type If interface {
	// InitRoutes initializes examples routes
	InitRoutes(r *chi.Mux)

	// getExample handler takes exampleID from url path variables and gets example.
	// Takes exampleID parameter from URL.
	// Returns:
	// [200] - if example received successfully.
	// [400] - if request or exampleID is malformed.
	// [401] - if unauthorized.
	// [404] - if example was not found.
	// [500] - if internal server issues occurred.
	getExample(w http.ResponseWriter, r *http.Request)
}

// Controller holds examples service
type Controller struct {
	examplesService examples_service.If
}

// Init creates Controller using service
func Init(examplesService examples_service.If) If {
	return &Controller{
		examplesService: examplesService,
	}
}

// InitRoutes initializes Examples routes.
func (c *Controller) InitRoutes(r *chi.Mux) {
	r.Get(constExampleURL, c.getExample)
}

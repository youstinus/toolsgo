package metrics

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// names used.
const (
	constMetricsURL = "/metrics"
)

// If used for handling metrics endpoints.
type If interface {
	InitRoutes(r *chi.Mux)
}

// Controller empty struct to use as metrics controller.
type Controller struct {
}

// Init creates metrics controller instance.
func Init() If {
	return &Controller{}
}

// InitRoutes initializes metrics handler.
func (c *Controller) InitRoutes(r *chi.Mux) {
	r.Handle(constMetricsURL, promhttp.Handler())
}

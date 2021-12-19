package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// prometheus query names
	reqsName    = "chi_requests_total"
	latencyName = "chi_request_duration_milliseconds"

	// used strings
	titleHistogramHelp = "How long it took to process the request, partitioned by status code, method and HTTP path."
	titleCounterHelp   = "How many HTTP requests processed, partitioned by status code, method and HTTP path."
	numNano            = 1000000
)

var (
	labelNames = []string{"code", "method", "path"}
	dflBuckets = []float64{100, 500, 1000}
)

// Middleware is a handler that exposes prometheus metrics for the number of requests,
// the latency and the response size, partitioned by status code, method and HTTP path.
type Middleware struct {
	reqs    *prometheus.CounterVec
	latency *prometheus.HistogramVec
}

// MetricsMiddleware returns a prometheus Middleware handler.
func MetricsMiddleware(buckets ...float64) func(_ http.Handler) http.Handler {
	var m Middleware
	m.reqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: reqsName,
			Help: titleCounterHelp,
		},
		labelNames,
	)
	prometheus.MustRegister(m.reqs)

	if len(buckets) == 0 {
		buckets = dflBuckets
	}

	m.latency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    latencyName,
			Help:    titleHistogramHelp,
			Buckets: buckets,
		},
		labelNames,
	)
	prometheus.MustRegister(m.latency)

	return m.handler
}

// handler prometheus metrics handler.
func (c Middleware) handler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r)
			c.reqs.WithLabelValues(http.StatusText(ww.Status()), r.Method, r.URL.Path).Inc()
			c.latency.WithLabelValues(http.StatusText(ww.Status()), r.Method, r.URL.Path).Observe(float64(time.Since(start).Nanoseconds()) / numNano)
		},
	)
}

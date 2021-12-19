package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	servermw "github.com/youstinus/toolsgo/api/server/middleware"
	"github.com/youstinus/toolsgo/api/server/router"
	"github.com/youstinus/toolsgo/api/server/router/metrics"
	"github.com/youstinus/toolsgo/pkg/types"
)

// Names used in package
const (
	emptyString   = ""
	titleStarting = "starting http server..."
	titleStopped  = "http server stopped"
	fieldAddress  = "address"

	numIdleTimeout   = 30
	numTimeout       = 60
	numHeaderTimeout = 60
	numWriteTimeout  = 60
	numCloseTimeout  = 10
)

// Options Public options
type Options struct {
	Enabled       bool
	Log           *logrus.Logger
	HTTPAddress   string
	HTTPPort      int
	EnableCors    bool
	AuthKey       string
	EnableMetrics bool
	SignalChan    chan os.Signal

	Controllers router.If // Controllers with services
}

// Server with private options
type Server struct {
	enabled       bool
	log           *logrus.Logger
	authKey       string
	httpServer    *http.Server
	httpAddr      string
	router        *chi.Mux
	started       bool
	enableCors    bool
	enableMetrics bool
	signalChan    chan os.Signal

	controllers router.If // controllers instance
}

// New Creates new private server with public options
func New(opts *Options) *Server {
	return &Server{
		enabled:       opts.Enabled,
		log:           opts.Log,
		authKey:       opts.AuthKey,
		httpAddr:      fmt.Sprintf("%s:%d", opts.HTTPAddress, opts.HTTPPort),
		enableCors:    opts.EnableCors,
		enableMetrics: opts.EnableMetrics,
		signalChan:    opts.SignalChan,
		controllers:   opts.Controllers,
	}
}

// ServeHTTP Starts Server after configuring
func (s *Server) ServeHTTP() {
	// do not start server unless enabled
	if !s.enabled {
		return
	}

	// configuring middleware, router
	s.configureServer(s.httpAddr)

	// Starts server in new go routine
	go func() {
		s.started = true
		s.log.WithField(fieldAddress, s.httpAddr).Info(titleStarting)

		defer s.log.Info(titleStopped)

		// Server holds here and server works until error occurs, then returns error and finishes program
		err := s.httpServer.ListenAndServe()
		s.started = false

		if err != nil { // if error is nil so app will not terminate
			s.log.Errorf(types.ErrfDump, err)
			s.signalChan <- syscall.SIGTERM
		}
	}()
}

// configureServer Configures server using properties, options. Sets handlers.
func (s *Server) configureServer(httpAddr string) {
	r := chi.NewRouter()
	s.router = r

	r.Use(middleware.RequestID)
	// changes server logger with logrus.
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  s.log,
		NoColor: true,
	}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	if s.enableCors {
		r.Use(servermw.CorsMiddleware())
	}

	// adds auth middleware if auth key setup in configuration.
	if s.authKey != emptyString {
		r.Use(servermw.AuthKeyMiddleware(s.authKey))
	}

	// enable metrics middleware for prometheus if enabled.
	if s.enableMetrics {
		r.Use(servermw.MetricsMiddleware())
		metrics.Init().InitRoutes(r)
	}

	// adds all routes to server using router.
	s.controllers.InitRoutes(r)

	// HTTP server
	s.httpServer = &http.Server{
		Addr:              httpAddr,
		Handler:           r,
		IdleTimeout:       time.Second * numIdleTimeout,
		ReadTimeout:       time.Second * numTimeout,
		ReadHeaderTimeout: time.Second * numHeaderTimeout,
		WriteTimeout:      time.Second * numWriteTimeout,
	}
}

// Close Manual server stopping.
func (s *Server) Close() {
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*numCloseTimeout)
		defer cancel()
		// it is not clear when Shutdown returns error so it is not handled.
		s.httpServer.Shutdown(ctx)
	}
}

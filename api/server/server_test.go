package server

import (
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/youstinus/toolsgo/api/server/router"
	"github.com/youstinus/toolsgo/api/server/router/examples"
)

func TestNew(t *testing.T) {
	sgn := make(chan os.Signal)
	lg := &logrus.Logger{}
	ctrls := &router.Controllers{}
	opts := &Options{
		Log:         lg,
		HTTPAddress: "localhost",
		HTTPPort:    8081,
		EnableCors:  true,
		AuthKey:     "test",
		SignalChan:  sgn,
		Controllers: ctrls,
	}

	type args struct {
		opts *Options
	}

	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "success - creates server from options",
			args: args{
				opts: opts,
			},
			want: &Server{
				log:         lg,
				authKey:     "test",
				httpAddr:    "localhost:8081",
				enableCors:  true,
				signalChan:  sgn,
				controllers: ctrls,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ServeHTTP(t *testing.T) {
	ctrls := &router.Controllers{
		Controller: &examples.Controller{},
	}
	sgn := make(chan os.Signal)

	tests := []struct {
		name string
		s    *Server
	}{
		// 1. success
		{
			name: "success configures and runs through the method",
			s: &Server{
				enabled:       true,
				log:           &logrus.Logger{},
				authKey:       "test",
				httpServer:    &http.Server{},
				httpAddr:      ":8082",
				router:        &chi.Mux{},
				started:       false,
				enableCors:    true,
				enableMetrics: true,
				signalChan:    sgn,
				controllers:   ctrls,
			},
		},

		// 2. disabled
		{
			name: "success enabled false",
			s: &Server{
				enabled: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Starts a go routine.
			tt.s.ServeHTTP()
			// Causes ListenAndServe to return "server is closed error"
			tt.s.Close()
		})
	}
}

func TestServer_Close(t *testing.T) {
	tests := []struct {
		name    string
		s       *Server
		wantErr bool
	}{
		{
			name: "success - shuts down",
			s: &Server{
				httpServer: &http.Server{},
			},
			wantErr: false,
		},
		{
			name:    "success - server is nil",
			s:       &Server{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Close()
		})
	}
}

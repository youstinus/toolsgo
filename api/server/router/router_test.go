package router

import (
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/youstinus/toolsgo/api/server/router/examples"
	"github.com/youstinus/toolsgo/pkg/services"
	"github.com/youstinus/toolsgo/pkg/services/examples_service"
)

func TestInitControllers(t *testing.T) {
	eService := examples_service.ExamplesService{}

	type args struct {
		s services.If
	}

	tests := []struct {
		name string
		args args
		want If
	}{
		{
			name: "success pass through",
			args: args{
				s: &services.Services{
					ExamplesService: &eService,
				},
			},
			want: &Controllers{
				Controller: examples.Init(&eService),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitControllers(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitControllers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestControllers_InitExamplesRoutes(t *testing.T) {
	type args struct {
		r *chi.Mux
	}

	tests := []struct {
		name string
		c    *Controllers
		args args
	}{
		{
			name: "success - calls all methods",
			args: args{
				r: chi.NewRouter(),
			},
			c: &Controllers{
				Controller: &examples.Controller{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.InitExamplesRoutes(tt.args.r)
		})
	}
}

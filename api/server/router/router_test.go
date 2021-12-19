package router

import (
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/youstinus/toolsgo/api/server/router/examples"
	"github.com/youstinus/toolsgo/api/server/router/tools"
	"github.com/youstinus/toolsgo/pkg/services"
	"github.com/youstinus/toolsgo/pkg/services/examples_service"
	"github.com/youstinus/toolsgo/pkg/services/toolsservice"
)

func TestInitControllers(t *testing.T) {
	eService := examples_service.ExamplesService{}
	tService := toolsservice.Service{}

	type args struct {
		s services.If
	}

	tests := []struct {
		name string
		args args
		want If
	}{
		{
			name: "Success",
			args: args{
				s: &services.Services{
					ExamplesService: &eService,
					ToolsService:    &tService,
				},
			},
			want: &Controllers{
				Controller: examples.Init(&eService),
				Tools:      tools.Init(&tService),
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

func TestControllers_InitRoutes(t *testing.T) {
	type args struct {
		r *chi.Mux
	}

	tests := []struct {
		name string
		c    *Controllers
		args args
	}{
		{
			name: "SuccessAllMethods",
			args: args{
				r: chi.NewRouter(),
			},
			c: &Controllers{
				Controller: &examples.Controller{},
				Tools:      &tools.Controller{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.InitRoutes(tt.args.r)
		})
	}
}

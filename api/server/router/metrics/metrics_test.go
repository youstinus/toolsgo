package metrics

import (
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		want If
	}{
		{
			name: "success pass through",
			want: &Controller{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_InitRoutes(t *testing.T) {
	type args struct {
		r *chi.Mux
	}

	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		{
			name: "success pass through",
			c:    &Controller{},
			args: args{
				r: chi.NewRouter(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.InitRoutes(tt.args.r)
		})
	}
}

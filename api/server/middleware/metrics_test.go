package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMetricsMiddleware(t *testing.T) {
	type args struct {
		buckets []float64
	}

	tests := []struct {
		name string
		args args
		want func(next http.Handler) http.Handler
	}{
		{
			name: "success pass through",
			args: args{},
			want: func(next http.Handler) http.Handler {
				return http.NotFoundHandler()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MetricsMiddleware(tt.args.buckets...)
			if got == nil {
				t.Errorf("MetricsMiddleware() = %v, want not nil", nil)
			}
			nextCalled := false
			// create a handler to use as "next" which will verify the request
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true
			})
			wh := got(nextHandler)
			req := httptest.NewRequest("GET", "http://testing", nil)
			rr := httptest.NewRecorder()
			// call the handler using a mock response recorder (we'll not use that anyway)
			wh.ServeHTTP(rr, req)
			// Check if next called
			if !nextCalled {
				t.Errorf("next not called")
			}
		})
	}
}

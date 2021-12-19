package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthKeyMiddleware(t *testing.T) {
	type args struct {
		key        string
		requestKey string
		addKey     bool
	}

	tests := []struct {
		name             string
		args             args
		statusExpected   int
		expectedNextCall bool
	}{
		{
			name: "success authorized with key",
			args: args{
				key:        "test",
				requestKey: "test",
				addKey:     true,
			},
			statusExpected:   200,
			expectedNextCall: true,
		},
		{
			name: "unauthorized wrong key",
			args: args{
				key:        "test",
				requestKey: "MALFORMED",
				addKey:     true,
			},
			statusExpected:   401,
			expectedNextCall: false,
		},
		{
			name: "key not provided",
			args: args{
				key:        "",
				requestKey: "",
				addKey:     false,
			},
			statusExpected:   200,
			expectedNextCall: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextCalled := false
			// create a handler to use as "next" which will verify the request
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true
			})
			// create the handler to test, using our custom "next" handler
			handlerToTest := AuthKeyMiddleware(tt.args.key)
			hndl := handlerToTest(nextHandler)
			// create a mock request to use
			req := httptest.NewRequest("GET", "http://testing", nil)
			if tt.args.addKey {
				req.Header.Add("X-Auth-Key", tt.args.requestKey)
			}
			rr := httptest.NewRecorder()
			// call the handler using a mock response recorder (we'll not use that anyway)
			hndl.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.statusExpected {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.statusExpected)
			}
			if nextCalled != tt.expectedNextCall {
				t.Errorf("next was expected to be called/not called: %v, expected call: %v", nextCalled, tt.expectedNextCall)
			}
		})
	}
}

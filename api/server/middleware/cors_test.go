package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorsMiddleware(t *testing.T) {
	nextCalled := false
	// create a handler to use as "next" which will verify the request
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
	})
	// create the handler to test, using our custom "next" handler
	handlerToTest := CorsMiddleware()
	hndl := handlerToTest(nextHandler)
	// create a mock request to use
	req := httptest.NewRequest("GET", "http://testing", nil)
	rr := httptest.NewRecorder()
	// call the handler using a mock response recorder (we'll not use that anyway)
	hndl.ServeHTTP(rr, req)

	if !nextCalled {
		t.Errorf("next not called")
	}
}

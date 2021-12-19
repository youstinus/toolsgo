package examples

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/youstinus/toolsgo/api/server/httputils"
	"github.com/youstinus/toolsgo/pkg/types"
)

const (
	numBase    = 10
	numBitSize = 64
)

// getExample handler takes ExampleID from url path variables and gets current Example.
// Takes ExampleID parameter from URL.
// Returns:
// [200] - if Example received successfully.
// [400] - if request or ExampleID is malformed.
// [401] - if unauthorized.
// [404] - if Example was not found.
// [500] - if internal server issues occurred.
func (c *Controller) getExample(w http.ResponseWriter, r *http.Request) {
	// Parses url param
	exampleStr := chi.URLParam(r, constExampleID)
	exampleID, err := strconv.ParseInt(exampleStr, numBase, numBitSize)

	if err != nil {
		render.Render(w, r, httputils.ErrBadRequest(err))
		return
	}
	// Gets example from examples service -> from database
	example, err := c.examplesService.GetExample(exampleID)
	if err != nil {
		render.Render(w, r, httputils.ErrNotFound(err))
		return
	}

	// Forms response body
	exampleBody := &ExampleBody{example}

	render.Status(r, http.StatusOK)
	render.Render(w, r, exampleBody)
}

// ExampleBody contains Example struct ref.
type ExampleBody struct {
	*types.Example
}

// Bind used for payload to bind from json.
func (p *ExampleBody) Bind(_ *http.Request) error {
	return nil
}

// Render transforms to json and checks if payload is empty, then creates empty payload if it is empty.
func (p *ExampleBody) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

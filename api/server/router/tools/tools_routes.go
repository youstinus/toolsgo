package tools

import (
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/youstinus/toolsgo/api/server/httputils"
)

func (c *Controller) encodeBase64(w http.ResponseWriter, r *http.Request) {
	input := DataBody{}
	if err := render.Bind(r, &input); err != nil {
		if err == io.EOF {
			render.Render(w, r, httputils.ErrMissingBody)
		} else {
			render.Render(w, r, httputils.ErrBadRequest(err))
		}

		return
	}
	// Encodes string to base64 string.
	encoded := c.toolsService.EncodeBase64(input.Data)
	// Forms response body.
	output := &DataBody{encoded}

	render.Status(r, http.StatusOK)
	render.Render(w, r, output)
}

// decodeBase64 decodes data given in body.
func (c *Controller) decodeBase64(w http.ResponseWriter, r *http.Request) {
	input := DataBody{}
	if err := render.Bind(r, &input); err != nil {
		if err == io.EOF {
			render.Render(w, r, httputils.ErrMissingBody)
		} else {
			render.Render(w, r, httputils.ErrBadRequest(err))
		}

		return
	}
	// Decodes base64 string to string or returns error.
	encoded, err := c.toolsService.DecodeBase64(input.Data)
	if err != nil {
		render.Render(w, r, httputils.ErrBadRequest(err))
		return
	}
	// Forms response body.
	output := &DataBody{encoded}

	render.Status(r, http.StatusOK)
	render.Render(w, r, output)
}

// DataBody contains Example struct ref.
type DataBody struct {
	Data string `json:"data"`
}

// Bind used for payload to bind from json.
func (p *DataBody) Bind(_ *http.Request) error {
	return nil
}

// Render transforms to json and checks if payload is empty, then creates empty payload if it is empty.
func (p *DataBody) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

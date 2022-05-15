package binder

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/temphia/temphia/backend/server/btypes/rtypes"
)

func (b *Binder) httpCall(request rtypes.HTTPRequest) *rtypes.HTTPResponse {
	resp, err := httpRequest(
		request.Path,
		request.Method,
		request.Headers,
		getBodyReader(request.Headers, request.Body),
	)

	if err != nil {
		return &rtypes.HTTPResponse{
			SatusCode: 500,
			Headers:   map[string][]string{"binding_intercepted_err": {"true"}},
			Body:      []byte(err.Error()),
		}
	}
	return resp
}

// private

func getBodyReader(header map[string]string, body []byte) io.Reader {
	contentType, ok := header["Content-Type"]

	if !ok && body != nil {
		return bytes.NewReader(body)
	}
	switch contentType {
	// fixme
	case "application/x-www-form-urlencoded":
	case "multipart/form-data":
	default:
		bytes.NewReader(body)
	}
	return nil
}

func httpRequest(u, method string, header map[string]string, body io.Reader) (*rtypes.HTTPResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	for headkey, headval := range header {
		req.Header.Set(headkey, headval)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &rtypes.HTTPResponse{
		Headers:   res.Header,
		Body:      bytes,
		SatusCode: res.StatusCode,
	}
	return r, nil
}

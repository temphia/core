package binder

import (
	"errors"
	"net/http"

	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/lib/kosher"
)

func (b *Binder) HTTPCall(request rtypes.HTTPRequest) *rtypes.HTTPResponse {
	return b.httpCall(request)
}

func (b *Binder) HTTPCacheFile(url string, headers map[string]string) error {
	return nil
}

func (b *Binder) QuickJsonGet(url string) ([]byte, error) {

	resp := b.httpCall(rtypes.HTTPRequest{
		Method:  http.MethodGet,
		Path:    url,
		Headers: map[string]string{},
		Body:    nil,
	})

	if resp.SatusCode != 200 {
		return nil, errors.New(kosher.Str(resp.Body))
	}
	return resp.Body, nil
}

func (b *Binder) QuickJsonPost(url string, data []byte) ([]byte, error) {
	resp := b.httpCall(rtypes.HTTPRequest{
		Method:  http.MethodGet,
		Path:    url,
		Headers: map[string]string{},
		Body:    data,
	})

	if resp.SatusCode != 200 {
		return nil, errors.New(kosher.Str(resp.Body))
	}
	return resp.Body, nil
}

package responder

import (
	"io"
	"net/http"
)

type Responder interface {
	Respond(r *http.Request, rw http.ResponseWriter)
}

type URLRedirect string

func (u URLRedirect) Respond(r *http.Request, rw http.ResponseWriter) {
	http.Redirect(rw, r, string(u), 301)
}

type Binary []byte

func (b Binary) Respond(r *http.Request, rw http.ResponseWriter) {

}

type stream struct {
	inner io.ReadCloser
}

func (s stream) Respond(r *http.Request, rw http.ResponseWriter) {
	io.Copy(rw, s.inner)
	s.inner.Close()
}

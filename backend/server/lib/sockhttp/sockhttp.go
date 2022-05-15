package sockhttp

import (
	"bytes"
	"context"
	"net"
	"net/http"
)

// SockHTTP is HTTP client through unix socket
type SockHTTP struct {
	client http.Client
}

func New(filepath string) *SockHTTP {
	return &SockHTTP{
		client: http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", filepath)
				},
			},
		},
	}
}

func (s *SockHTTP) PostWithHeader(path string, payload []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, "http://unix"+path, bytes.NewReader(payload))

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if err != nil {
		return nil, err
	}
	return s.client.Do(req)
}

func (s *SockHTTP) Get(path string) (*http.Response, error) {
	return s.client.Get("http://unix" + path)
}

func (s *SockHTTP) Post(path string, payload []byte) (*http.Response, error) {
	return s.client.Post("http://unix"+path, "application/octet-stream", bytes.NewReader(payload))
}

func (s *SockHTTP) PostJSON(path string, payload []byte) (*http.Response, error) {
	return s.client.Post("http://unix"+path, "application/json", bytes.NewReader(payload))
}

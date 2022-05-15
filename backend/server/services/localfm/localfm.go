package localfm

import (
	"io"
	"sync"
)

type LocalFM struct {
	cache map[string]*cache
	mlock sync.Mutex
}

type cache struct {
	files map[string]string
}

func (l *LocalFM) PutFile(tenantId string, key string, data []byte) error {
	return nil
}

func (l *LocalFM) PutFileReader(tenantId string, key string, r io.ReadCloser) error {
	return nil
}

func (l *LocalFM) GetFile(tenantId string, key string) ([]byte, error) {
	return nil, nil
}

func (l *LocalFM) GetFileReader(tenantId string, r io.ReadCloser) ([]byte, error) {
	return nil, nil
}

func (l *LocalFM) RemoveFile(tenantId string, key ...string) error {
	return nil
}

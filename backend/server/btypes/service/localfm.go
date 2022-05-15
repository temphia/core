package service

import "io"

type LocalFileManager interface {
	PutFile(tenantId string, key string, data []byte) error
	PutFileReader(tenantId string, key string, r io.ReadCloser) error
	GetFile(tenantId string, key string) ([]byte, error)
	GetFileReader(tenantId string, r io.ReadCloser) ([]byte, error)
	RemoveFile(tenantId string, key ...string) error
}

package store

import (
	"github.com/temphia/core/backend/server/app/config"
	"github.com/temphia/core/backend/server/btypes/models/entities"
)

type BPrintLocalStore interface {
	BprintPut(tenantId string, et *entities.BPrint) (string, error)
	BprintGet(tenantId, id string) (*entities.BPrint, error)
	BprintDel(tenantId, id string) error
	BprintList(tenantId, group string) ([]*entities.BPrint, error)
}

type RepoBuilderOptions struct {
	TenantId      string
	BaseURL       string
	SourceOptions *config.StoreSource
}

type RepoBuilder func(opts *RepoBuilderOptions) (Repository, error)

type RepoQuery struct {
	Group string
	Tags  []string
	Page  int64
}

type Repository interface {
	Name() string
	Query(tenantId string, opts *RepoQuery) ([]entities.BPrint, error)
	GetItem(tenantid, group, slug string) (*entities.BPrint, error)
	GetFile(tenantid, group, slug, file string) ([]byte, error)
	GetFileURL(tenantid, group, slug, file string) (string, error)
}

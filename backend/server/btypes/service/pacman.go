package service

import (
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/models/vmodels"
	"gitlab.com/mr_balloon/golib/hmap"
)

type Pacman interface {
	// repo
	RepoSourceList(tenantid string) (map[string]string, error)
	RepoList(tenantid, source, group string, tags ...string) ([]entities.BPrint, error)
	RepoGet(tenantid, source, group, slug string) (*entities.BPrint, error)
	RepoImport(tenantid, source, group, slug string, data hmap.H) (string, error)
	RepoImportAndInstall(tenantid, source string, data hmap.H)
	RepoGetBlob(tenantid, source, group, slug, file string) ([]byte, error)

	// bprint

	BprintList(tenantid, group string) ([]*entities.BPrint, error)
	BprintCreate(tenantid string, bp *entities.BPrint) (string, error)
	BprintUpdate(tenantid string, bp *entities.BPrint) error
	BprintGet(tenantid, bid string) (*entities.BPrint, error)
	BprintRemove(tenantid, bid string) error
	BprintListBlobs(tenantid, bid string) (interface{}, error)

	BprintNewBlob(tenantid, bid, file string, payload []byte) error
	BprintUpdateBlob(tenantid, bid, file string, payload []byte) error

	BprintGetBlob(tenantid, bid, file string) ([]byte, error)
	BprintDeleteBlob(tenantid, bid, file string) error
	Install(tenantId string, opts *vmodels.RepoInstallOpts) (interface{}, error)
}

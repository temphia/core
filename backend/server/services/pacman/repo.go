package pacman

import (
	"context"

	"github.com/rs/xid"
	"github.com/temphia/temphia/backend/server/btypes"
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/store"
	"gitlab.com/mr_balloon/golib/hmap"
)

func (b *PacMan) RepoSourceList(tenantid string) (map[string]string, error) {
	resp := make(map[string]string)
	for srcname, src := range b.repos {
		resp[srcname] = src.Name()
	}

	return resp, nil
}

func (b *PacMan) RepoList(tenantid, source, group string, tags ...string) ([]entities.BPrint, error) {
	sp, err := b.repo(source)
	if err != nil {
		return nil, err
	}
	return sp.Query(tenantid, &store.RepoQuery{
		Group: group,
		Tags:  tags,
	})
}

func (b *PacMan) RepoGet(tenantid, source, group, slug string) (*entities.BPrint, error) {
	sp, err := b.repo(source)
	if err != nil {
		return nil, err
	}
	return sp.GetItem(tenantid, group, slug)
}

func (b *PacMan) RepoImport(tenantid, source, group, slug string, data hmap.H) (string, error) {
	repo, err := b.repo(source)
	if err != nil {
		return "", err
	}

	bp, err := repo.GetItem(tenantid, group, slug)
	if err != nil {
		return "", err
	}
	bp.Combine(data)

	bp.ID = xid.New().String()
	bp.TenantID = tenantid

	_, err = b.BprintCreate(tenantid, bp)
	if err != nil {
		return "", err
	}

	bstore := b.blobStore(tenantid)

	for _, fkey := range bp.Files {
		bytes, err := repo.GetFile(tenantid, group, slug, fkey)
		if err != nil {
			return "", err
		}
		err = bstore.AddBlob(context.TODO(), btypes.BprintBlobFolder, ffile(bp.ID, fkey), bytes)
		if err != nil {
			return "", err
		}

	}

	return bp.ID, nil
}

func (b *PacMan) RepoImportAndInstall(tenantid, source string, data hmap.H) {

}

func (b *PacMan) RepoGetBlob(tenantid, source, group, slug, file string) ([]byte, error) {
	sp, err := b.repo(source)
	if err != nil {
		return nil, err
	}
	return sp.GetFile(tenantid, group, slug, file)
}

func (b *PacMan) repo(src string) (store.Repository, error) {
	provider, ok := b.repos[src]
	if !ok {
		return nil, easyerr.NotFound()
	}

	return provider, nil
}

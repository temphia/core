package basic

import (
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
)

type RepoListReq struct {
	Source string
	Group  string
	Tags   []string
}

type RepoGetReq struct {
	Source string
	Group  string
	Slug   string
}

type RepoGetBlobReq struct {
	Source string
	Group  string
	Slug   string
	File   string
}

func (c *Controller) RepoSources(uclaim *claim.Session) (map[string]string, error) {
	return c.pacman.RepoSourceList(uclaim.TenentID)
}

func (c *Controller) RepoList(uclaim *claim.Session, req RepoListReq) ([]entities.BPrint, error) {
	return c.pacman.RepoList(uclaim.TenentID, req.Source, req.Group)
}

func (c *Controller) RepoGet(uclaim *claim.Session, req RepoGetReq) (*entities.BPrint, error) {
	return c.pacman.RepoGet(uclaim.TenentID, req.Source, req.Group, req.Slug)
}
func (c *Controller) RepoGetBlob(uclaim *claim.Session, req RepoGetBlobReq) ([]byte, error) {
	return c.pacman.RepoGetBlob(uclaim.TenentID, req.Source, req.Group, req.Slug, req.File)
}

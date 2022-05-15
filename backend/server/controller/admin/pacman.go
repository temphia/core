package admin

import (
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/models/vmodels"
)

func (c *Controller) BprintList(uclaim *claim.Session, group string) ([]*entities.BPrint, error) {

	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintList(uclaim.TenentID, group)
}

func (c *Controller) BprintCreate(uclaim *claim.Session, bp *entities.BPrint) (string, error) {
	if !uclaim.IsSuperAdmin() {
		return "", easyerr.NotImpl()
	}

	return c.pacman.BprintCreate(uclaim.TenentID, bp)
}

func (c *Controller) BprintUpdate(uclaim *claim.Session, bp *entities.BPrint) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintUpdate(uclaim.TenentID, bp)
}

func (c *Controller) BprintGet(uclaim *claim.Session, bid string) (*entities.BPrint, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintGet(uclaim.TenentID, bid)
}

func (c *Controller) BprintRemove(uclaim *claim.Session, bid string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintRemove(uclaim.TenentID, bid)
}

func (c *Controller) BprintListBlobs(uclaim *claim.Session, bid string) (interface{}, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintListBlobs(uclaim.TenentID, bid)
}

func (c *Controller) BprintNewBlob(uclaim *claim.Session, bid, file string, payload []byte) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintNewBlob(uclaim.TenentID, bid, file, payload)
}

func (c *Controller) BprintUpdateBlob(uclaim *claim.Session, bid, file string, payload []byte) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}
	return c.pacman.BprintUpdateBlob(uclaim.TenentID, bid, file, payload)
}

func (c *Controller) BprintGetBlob(uclaim *claim.Session, bid, file string) ([]byte, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintGetBlob(uclaim.TenentID, bid, file)
}

func (c *Controller) BprintDeleteBlob(uclaim *claim.Session, bid, file string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintDeleteBlob(uclaim.TenentID, bid, file)
}

// repo

func (c *Controller) BprintImport(uclaim *claim.Session, opts *vmodels.RepoImportOpts) (string, error) {
	return c.pacman.RepoImport(uclaim.TenentID, opts.Source, opts.Group, opts.Slug, opts.Options)
}

func (c *Controller) BprintInstall(uclaim *claim.Session, opts *vmodels.RepoInstallOpts) (interface{}, error) {
	opts.UserId = uclaim.UserID
	return c.pacman.Install(uclaim.TenentID, opts)
}

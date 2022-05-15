package admin

import (
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
)

func (c *Controller) AddUser(uclaim *claim.Session, user *entities.User) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.AddUser(user)
}

func (c *Controller) UpdateUser(uclaim *claim.Session, user map[string]interface{}) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.UpdateUser(uclaim.TenentID, uclaim.UserID, user)
}

func (c *Controller) RemoveUser(uclaim *claim.Session, username string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.RemoveUser(uclaim.TenentID, username)
}

func (c *Controller) GetUserByID(uclaim *claim.Session, username string) (*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.GetUserByID(uclaim.TenentID, username)
}

func (c *Controller) GetUserByEmail(uclaim *claim.Session, email string) (*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.GetUserByEmail(uclaim.TenentID, email)
}

func (c *Controller) ListUsers(uclaim *claim.Session) ([]*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.ListUsers(uclaim.TenentID)
}

func (c *Controller) ListUsersByGroup(uclaim *claim.Session, group string) ([]*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.ListUsersByGroup(uclaim.TenentID, group)
}

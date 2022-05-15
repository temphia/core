package admin

import (
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
)

func (c *Controller) AddUserGroup(uclaim *claim.Session, ugroup *entities.UserGroup) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.AddUserGroup(ugroup)
}

func (c *Controller) ListUserGroup(uclaim *claim.Session) ([]*entities.UserGroup, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.ListUserGroups(uclaim.TenentID)
}

func (c *Controller) GetUserGroup(uclaim *claim.Session, ugroup string) (*entities.UserGroup, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.GetUserGroup(uclaim.TenentID, ugroup)
}

func (c *Controller) UpdateUserGroup(uclaim *claim.Session, id string, data map[string]interface{}) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.UpdateUserGroup(uclaim.TenentID, id, data)
}

func (c *Controller) RemoveUserGroup(uclaim *claim.Session, ugroup string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.RemoveUserGroup(uclaim.TenentID, ugroup)
}

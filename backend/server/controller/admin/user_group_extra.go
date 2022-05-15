package admin

import (
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/models/entities"
)

// auth

func (c *Controller) AddUserGroupAuth(uclaim *claim.Session, gslug string, data *entities.UserGroupAuth) error {
	data.TenantId = uclaim.TenentID
	data.UserGroup = gslug
	return c.coredb.AddUserGroupAuth(data)
}

func (c *Controller) UpdateUserGroupAuth(uclaim *claim.Session, gslug string, id int64, data map[string]interface{}) error {
	return c.coredb.UpdateUserGroupAuth(uclaim.TenentID, gslug, id, data)
}

func (c *Controller) ListUserGroupAuth(uclaim *claim.Session, gslug string) ([]*entities.UserGroupAuth, error) {
	return c.coredb.ListUserGroupAuth(uclaim.TenentID, gslug)
}

func (c *Controller) GetUserGroupAuth(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupAuth, error) {
	return c.coredb.GetUserGroupAuth(uclaim.TenentID, gslug, id)
}

func (c *Controller) RemoveUserGroupAuth(uclaim *claim.Session, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupAuth(uclaim.TenentID, gslug, id)
}

// hook

func (c *Controller) AddUserGroupHook(uclaim *claim.Session, gslug string, data *entities.UserGroupHook) error {
	data.TenantId = uclaim.TenentID
	data.UserGroup = gslug
	return c.coredb.AddUserGroupHook(data)
}

func (c *Controller) UpdateUserGroupHook(uclaim *claim.Session, gslug string, id int64, data map[string]interface{}) error {
	return c.coredb.UpdateUserGroupHook(uclaim.TenentID, gslug, id, data)
}

func (c *Controller) ListUserGroupHook(uclaim *claim.Session, gslug string) ([]*entities.UserGroupHook, error) {
	return c.coredb.ListUserGroupHook(uclaim.TenentID, gslug)
}

func (c *Controller) GetUserGroupHook(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupHook, error) {
	return c.coredb.GetUserGroupHook(uclaim.TenentID, gslug, id)
}

func (c *Controller) RemoveUserGroupHook(uclaim *claim.Session, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupHook(uclaim.TenentID, gslug, id)
}

// plug

func (c *Controller) AddUserGroupPlug(uclaim *claim.Session, gslug string, data *entities.UserGroupPlug) error {
	data.TenantId = uclaim.TenentID
	data.UserGroup = gslug
	return c.coredb.AddUserGroupPlug(data)
}

func (c *Controller) UpdateUserGroupPlug(uclaim *claim.Session, gslug string, id int64, data map[string]interface{}) error {
	return c.coredb.UpdateUserGroupPlug(uclaim.TenentID, gslug, id, data)
}

func (c *Controller) ListUserGroupPlug(uclaim *claim.Session, gslug string) ([]*entities.UserGroupPlug, error) {
	return c.coredb.ListUserGroupPlug(uclaim.TenentID, gslug)
}

func (c *Controller) GetUserGroupPlug(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupPlug, error) {
	return c.coredb.GetUserGroupPlug(uclaim.TenentID, gslug, id)
}

func (c *Controller) RemoveUserGroupPlug(uclaim *claim.Session, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupPlug(uclaim.TenentID, gslug, id)
}

// data

func (c *Controller) AddUserGroupData(uclaim *claim.Session, gslug string, data *entities.UserGroupData) error {
	data.TenantId = uclaim.TenentID
	data.UserGroup = gslug
	return c.coredb.AddUserGroupData(data)
}

func (c *Controller) UpdateUserGroupData(uclaim *claim.Session, gslug string, id int64, data map[string]interface{}) error {
	return c.coredb.UpdateUserGroupData(uclaim.TenentID, gslug, id, data)
}

func (c *Controller) ListUserGroupData(uclaim *claim.Session, gslug string) ([]*entities.UserGroupData, error) {
	return c.coredb.ListUserGroupData(uclaim.TenentID, gslug)
}

func (c *Controller) GetUserGroupData(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupData, error) {
	return c.coredb.GetUserGroupData(uclaim.TenentID, gslug, id)
}

func (c *Controller) RemoveUserGroupData(uclaim *claim.Session, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupData(uclaim.TenentID, gslug, id)
}

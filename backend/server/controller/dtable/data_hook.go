package dtable

import (
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/models/entities"
)

func (c *Controller) NewHook(uclaim *claim.Session, tslug string, model *entities.DataHook) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	model.GroupID = uclaim.ServicePath[2]
	model.TableID = tslug
	model.TenantID = uclaim.TenentID
	return dynDB.NewHook(model)
}

func (c *Controller) ModifyHook(uclaim *claim.Session, tslug string, id int64, data map[string]interface{}) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return dynDB.ModifyHook(uclaim.ServicePath[2], tslug, id, data)
}

func (c *Controller) ListHook(uclaim *claim.Session, tslug string) ([]*entities.DataHook, error) {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return dynDB.ListHook(uclaim.ServicePath[2], tslug)
}

func (c *Controller) DelHook(uclaim *claim.Session, tslug string, id int64) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.DelHook(uclaim.ServicePath[2], tslug, id)
}

func (c *Controller) GetHook(uclaim *claim.Session, tslug string, id int64) (*entities.DataHook, error) {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.GetHook(uclaim.ServicePath[2], tslug, id)
}

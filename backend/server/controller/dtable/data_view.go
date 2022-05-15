package dtable

import (
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/models/entities"
)

func (c *Controller) NewView(uclaim *claim.Session, tslug string, model *entities.DataView) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	model.GroupID = uclaim.ServicePath[2]
	model.TableID = tslug
	model.TenantID = uclaim.TenentID

	return dynDB.NewView(model)
}

func (c *Controller) ModifyView(uclaim *claim.Session, tslug string, id int64, data map[string]interface{}) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return dynDB.ModifyView(uclaim.ServicePath[2], tslug, id, data)
}

func (c *Controller) ListView(uclaim *claim.Session, tslug string) ([]*entities.DataView, error) {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return dynDB.ListView(uclaim.ServicePath[2], tslug)
}

func (c *Controller) DelView(uclaim *claim.Session, tslug string, id int64) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return dynDB.DelView(uclaim.ServicePath[2], tslug, id)
}

func (c *Controller) GetView(uclaim *claim.Session, tslug string, id int64) (*entities.DataView, error) {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.GetView(uclaim.ServicePath[2], tslug, id)
}

package admin

import (
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/models/entities"
)

func (c *Controller) GetTenant(uclaim *claim.Session) (*entities.Tenant, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}
	return c.coredb.GetTenant(uclaim.TenentID)
}

func (c *Controller) UpdateTenant(uclaim *claim.Session, data map[string]interface{}) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}
	slug := data["slug"].(string)

	delete(data, "slug")

	return c.coredb.UpdateTenant(slug, data)
}

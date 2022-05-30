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

// domain

func (c *Controller) AddDomain(uclaim *claim.Session, domain *entities.TenantDomain) error {
	domain.TenantId = uclaim.TenentID
	return c.coredb.AddDomain(domain)
}

func (c *Controller) UpdateDomain(uclaim *claim.Session, id int64, data map[string]interface{}) error {
	return c.coredb.UpdateDomain(uclaim.TenentID, id, data)
}

func (c *Controller) GetDomain(uclaim *claim.Session, id int64) (*entities.TenantDomain, error) {
	return c.coredb.GetDomain(uclaim.TenentID, id)
}

func (c *Controller) RemoveDomain(uclaim *claim.Session, id int64) error {
	return c.coredb.RemoveDomain(uclaim.TenentID, id)
}

func (c *Controller) ListDomain(uclaim *claim.Session) ([]*entities.TenantDomain, error) {
	return c.coredb.ListDomain(uclaim.TenentID)
}

// widget

func (c *Controller) AddDomainWidget(uclaim *claim.Session, domain *entities.DomainWidget) error {
	domain.TenantId = uclaim.TenentID
	return c.coredb.AddDomainWidget(domain)
}

func (c *Controller) UpdateDomainWidget(uclaim *claim.Session, id int64, data map[string]interface{}) error {
	return c.coredb.UpdateDomainWidget(uclaim.TenentID, id, data)
}

func (c *Controller) GetDomainWidget(uclaim *claim.Session, id int64) (*entities.DomainWidget, error) {
	return c.coredb.GetDomainWidget(uclaim.TenentID, id)
}

func (c *Controller) RemoveDomainWidget(uclaim *claim.Session, id int64) error {
	return c.coredb.RemoveDomainWidget(uclaim.TenentID, id)
}

func (c *Controller) ListDomainWidget(uclaim *claim.Session, did int64) ([]*entities.DomainWidget, error) {
	return c.coredb.ListDomainWidget(uclaim.TenentID, did)
}

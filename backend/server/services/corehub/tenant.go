package corehub

import (
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/service"
)

func (c *CoreHub) AddTenant(tenant *entities.Tenant) error {
	err := c.coredb.AddTenant(tenant)
	if err != nil {
		return err
	}

	eb := c.cplane.GetEventBus()
	eb.EmitTenantEvent(tenant.Slug, service.EventCreateTenant, tenant)
	return nil
}

func (c *CoreHub) UpdateTenant(slug string, data map[string]interface{}) error {
	return c.coredb.UpdateTenant(slug, data)
}

func (c *CoreHub) GetTenant(tenant string) (*entities.Tenant, error) {
	return c.coredb.GetTenant(tenant)
}

func (c *CoreHub) RemoveTenant(slug string) error {
	return c.coredb.RemoveTenant(slug)
}

func (c *CoreHub) ListTenant() ([]*entities.Tenant, error) {
	return c.coredb.ListTenant()
}

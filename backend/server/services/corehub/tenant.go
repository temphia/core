package corehub

import "github.com/temphia/temphia/backend/server/btypes/models/entities"

func (c *CoreHub) AddTenant(tenant *entities.Tenant) error {
	return c.coredb.AddTenant(tenant)
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

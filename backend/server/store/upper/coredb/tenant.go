package coredb

import (
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) AddTenant(tenant *entities.Tenant) error {
	_, err := d.tenantTable().Insert(tenant)
	return err
}

func (d *DB) UpdateTenant(slug string, data map[string]interface{}) error {
	return d.tenantTable().Find(db.Cond{"slug": slug}).Update(data)
}

func (d *DB) GetTenant(slug string) (*entities.Tenant, error) {
	ten := &entities.Tenant{}
	err := d.tenantTable().Find(db.Cond{"slug": slug}).One(ten)
	if err != nil {
		return nil, err
	}
	return ten, nil
}

func (d *DB) RemoveTenant(slug string) error {
	return d.tenantTable().Find(db.Cond{"slug": slug}).Delete()
}

func (d *DB) ListTenant() ([]*entities.Tenant, error) {
	tens := make([]*entities.Tenant, 0)
	err := d.tenantTable().Find().All(&tens)
	if err != nil {
		return nil, err
	}

	return tens, nil
}

// private

func (d *DB) tenantTable() db.Collection {
	return d.table("tenants")
}

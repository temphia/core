package cabinet

import (
	"github.com/temphia/core/backend/server/app/config"
	"github.com/temphia/core/backend/server/btypes/store"
)

func NewAdapter(name string, provider func(*config.StoreSource) (store.CabinetSource, error)) func(*config.StoreSource) (store.Store, error) {
	return func(ss *config.StoreSource) (store.Store, error) {

		cs, err := provider(ss)
		if err != nil {
			return nil, err
		}

		return &CabinetAdaper{
			cabSource: cs,
			name:      name,
		}, nil
	}
}

type CabinetAdaper struct {
	cabSource store.CabinetSource
	name      string
}

func (c *CabinetAdaper) Supports(_type store.StoreType) bool {
	if _type == store.TypeBlobStore {
		return true
	}
	return false
}

func (c *CabinetAdaper) CabinetSource() store.CabinetSource { return c.cabSource }
func (c *CabinetAdaper) CoreDB() store.CoreDB               { return nil }
func (c *CabinetAdaper) DynDB() store.DynDB                 { return nil }
func (c *CabinetAdaper) StateDB() store.PlugStateKV         { return nil }
func (c *CabinetAdaper) SyncDB() store.SyncDB               { return nil }
func (c *CabinetAdaper) Name() string                       { return c.name }

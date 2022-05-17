package cabinethub

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/btypes/store"
)

type CabinetHub struct {
	defaultProvider store.CabinetSource
	sources         map[string]store.CabinetSource
	defName         string
}

var defaultFolders = []string{"bprints", "data_common", "public"}

func New(sources map[string]store.CabinetSource, defprovider string) *CabinetHub {
	ch := &CabinetHub{
		sources:         sources,
		defaultProvider: sources[defprovider],
		defName:         defprovider,
	}
	return ch
}

func (c *CabinetHub) Start(eventbus interface{}) error {
	eb := eventbus.(service.EventBus)

	eb.OnTenantChange(func(tenant, event string, data *entities.Tenant) {
		go func() {
			switch event {
			case service.EventCreateTenant:
				c.defaultProvider.InitilizeTenent(tenant, defaultFolders)
			default:
				pp.Println("skipping event")
			}
		}()
	})

	return nil
}

func (c *CabinetHub) Default(tenant string) store.CabinetSourced {
	return &cabinetSourced{
		source:   "default",
		tenantId: tenant,
		provider: c.defaultProvider,
	}
}

func (c *CabinetHub) ListSources(tenant string) ([]string, error) {
	sources := make([]string, 0, len(c.sources))
	for k := range c.sources {
		sources = append(sources, k)
	}
	return sources, nil
}

func (c *CabinetHub) GetSource(source, tenant string) store.CabinetSourced {

	pp.Println(source, tenant)

	provider, ok := c.sources[source]
	if !ok {
		panic(easyerr.NotFound())
	}

	return &cabinetSourced{
		source:   source,
		tenantId: tenant,
		provider: provider,
	}
}

func (c *CabinetHub) DefaultName(tenant string) string { return c.defName }

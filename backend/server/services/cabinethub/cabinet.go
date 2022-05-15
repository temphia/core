package cabinethub

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/store"
)

type CabinetHub struct {
	defaultProvider store.CabinetSource
	sources         map[string]store.CabinetSource
	defName         string
}

func New(sources map[string]store.CabinetSource, defprovider string) *CabinetHub {
	return &CabinetHub{
		sources:         sources,
		defaultProvider: sources[defprovider],
		defName:         defprovider,
	}

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

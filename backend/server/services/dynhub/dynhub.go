package dynhub

import (
	"github.com/temphia/temphia/backend/server/btypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/temphia/temphia/backend/server/btypes/store"
	"github.com/temphia/temphia/backend/server/services/sockdhub"
)

type DynHub struct {
	dyndbs   map[string]store.DynDB
	eventHub service.EventBus
	sockdhub sockdhub.SockdHub
	engine   rtypes.Engine
}

func New(_app btypes.App, dyns map[string]store.DynDB) *DynHub {
	return &DynHub{
		dyndbs:   dyns,
		eventHub: _app.ControlPlane().(service.ControlPlane).GetEventBus(),
		sockdhub: *sockdhub.New(_app.Sockd().(service.SockCore)),
	}
}

func (s *DynHub) GetSource(source, tenant string) store.DynSource {
	return &dynSource{
		hub:      s,
		source:   source,
		tenantId: tenant,
	}
}

func (s *DynHub) ListSources(tenantId string) ([]string, error) {

	sources := make([]string, 0, len(s.dyndbs))
	for srcName := range s.dyndbs {
		sources = append(sources, srcName)
	}

	return sources, nil
}

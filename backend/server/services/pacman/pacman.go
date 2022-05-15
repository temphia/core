package pacman

import (
	"github.com/temphia/temphia/backend/server/btypes"
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/temphia/temphia/backend/server/btypes/store"
)

var (
	_ service.Pacman = (*PacMan)(nil)
)

type PacMan struct {
	app     btypes.App
	sockd   service.SockCore
	syncer  store.SyncDB
	dynHub  store.DynHub
	repos   map[string]store.Repository
	cabinet store.CabinetHub
}

func New(_app btypes.App, repos map[string]store.Repository) *PacMan {

	return &PacMan{
		app:     _app,
		sockd:   _app.Sockd().(service.SockCore),
		syncer:  _app.CoreHub(),
		repos:   repos,
		dynHub:  _app.DynHub(),
		cabinet: _app.Cabinet(),
	}
}

func (p *PacMan) blobStore(tenantId string) store.CabinetSourced {
	return p.cabinet.Default(tenantId)
}

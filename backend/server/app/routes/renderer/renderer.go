package renderer

import (
	"sync"

	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/service"
)

type Renderer struct {
	renderers map[string]*DomainRenderer
	mlock     sync.Mutex
	engine    rtypes.Engine
	pacman    service.Pacman
}

func New(engine rtypes.Engine, pacman service.Pacman) *Renderer {
	return &Renderer{
		renderers: make(map[string]*DomainRenderer),
		mlock:     sync.Mutex{},
		engine:    engine,
		pacman:    pacman,
	}
}

func (r *Renderer) Render(domain string) ([]byte, error) {
	return r.render(domain)
}

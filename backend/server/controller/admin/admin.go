package admin

import (
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/temphia/temphia/backend/server/btypes/store"
)

type Controller struct {
	pacman service.Pacman
	cplane service.ControlPlane
	coredb store.CoreHub
	signer service.Signer
}

func New(pacman service.Pacman, cplane service.ControlPlane, coredb store.CoreHub, signer service.Signer) *Controller {
	return &Controller{
		pacman: pacman,
		cplane: cplane,
		coredb: coredb,
		signer: signer,
	}
}

package corehub

import (
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/temphia/temphia/backend/server/btypes/store"
	"github.com/temphia/temphia/backend/server/services/sockdhub"
)

type CoreHub struct {
	coredb   store.CoreDB
	sockdhub sockdhub.SockdHub
	cplane   service.ControlPlane
}

func New(coredb store.CoreDB, sockd service.SockCore, cplane service.ControlPlane) *CoreHub {
	return &CoreHub{
		coredb:   coredb,
		sockdhub: *sockdhub.New(sockd),
		cplane:   cplane,
	}
}

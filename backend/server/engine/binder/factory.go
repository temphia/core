package binder

import (
	"github.com/rs/zerolog"
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/service"
)

type Factory struct {
	App    btypes.App
	Sockd  service.SockCore
	Pacman service.Pacman
	Logger zerolog.Logger
}

type BinderOptions struct {
	Namespace string
	PlugId    string
}

func (bf Factory) New(opts BinderOptions) *Binder {

	b := &Binder{
		factory:   bf,
		namespace: opts.Namespace,
		plugId:    opts.PlugId,
		ipcBinder: ipcBinder{
			core:   nil,
			loaded: false,
			slots:  make(map[string]interface{}),
		},
		plugKvBinder: plugKvBinder{
			core:    nil,
			stateKv: nil,
			txns:    []uint32{},
		},
		resources: nil,
	}

	return b
}

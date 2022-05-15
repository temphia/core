package rtypes

import (
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/models/entities"
)

type ModuleOption struct {
	Binder   ExecutorBinder
	Resource *entities.Resource
}

type Module interface {
	IPC(method string, path string, args btypes.Data) (btypes.Data, error)
	Close() error
}

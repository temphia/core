package rtypes

import (
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/event"
)

// Executor stuff
type ExecutorOption struct {
	Binder   Bindings
	PlugId   string
	AgentId  string
	Slug     string
	ExecType string
}

type Executor interface {
	Process(*event.Request) (*event.Response, error)
}

type ExecutorBuilder interface {
	Instance(ExecutorOption) (Executor, error)
	ExecFile(file string) ([]byte, error)
	// Init(app interface{}) error
}

type BuilderFunc func(ExecutorOption) (Executor, error)

func (f BuilderFunc) Instance(opt ExecutorOption) (Executor, error) {
	return f(opt)
}

func (f BuilderFunc) ExecFile(file string) ([]byte, error) {
	return nil, easyerr.NotFound()
}

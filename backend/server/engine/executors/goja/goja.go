package goja

import (
	"errors"

	"github.com/dop251/goja"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/event"
)

var _ rtypes.Executor = (*Goja)(nil)

type Goja struct {
	runtime *goja.Runtime
	binder  rtypes.Bindings
}

func NewBuilder(opts rtypes.ExecutorOption) (rtypes.Executor, error) {
	script, err := opts.Binder.GetSelfFile("server.js")
	if err != nil {
		return nil, err
	}

	rt := goja.New()
	_, err = rt.RunString(string(script))
	if err != nil {
		return nil, err
	}

	return New(opts.Binder, rt)
}

func New(b rtypes.Bindings, rt *goja.Runtime) (*Goja, error) {

	g := &Goja{
		runtime: rt,
		binder:  b,
	}

	g.bind()

	return g, nil
}

func (g *Goja) Process(ev *event.Request) (*event.Response, error) {
	var entry func(ev *event.Request) (*event.Response, error)
	rawentry := g.runtime.Get(ev.Name)
	if rawentry == nil {
		return nil, errors.New("method not found")
	}

	err := g.runtime.ExportTo(rawentry, &entry)
	if err != nil {
		return nil, err
	}
	return entry(ev)
}

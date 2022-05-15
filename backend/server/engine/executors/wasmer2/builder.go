package wasmer2

import (
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func NewBuilder() func(opts rtypes.ExecutorOption) (rtypes.Executor, error) {
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	return func(opts rtypes.ExecutorOption) (rtypes.Executor, error) {
		out, err := opts.Binder.GetSelfFile("server.wasm")
		if err != nil {
			return nil, err
		}

		exec, err := New(store, opts.Binder, out)
		if err != nil {
			return nil, err
		}
		err = exec.init()
		if err != nil {
			return nil, err
		}
		return exec, nil
	}

}

func (w *wasmer2) Process(ereq *event.Request) (*event.Response, error) {

	entry, err := w.instance.Exports.GetFunction(ereq.Name)
	if err != nil {
		return nil, err
	}

	w.eventRequest = ereq

	_, err = entry()
	if err != nil {
		return nil, err
	}

	if w.eventReply == nil {
		return nil, easyerr.Error("Empty Reply")
	}

	return w.eventReply, nil
}

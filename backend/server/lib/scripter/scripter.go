package scripter

import (
	"github.com/dop251/goja"
	"github.com/temphia/core/backend/server/btypes/easyerr"
)

type ScriptExecutor struct {
	Runtime   *goja.Runtime
	BindFuncs map[string]interface{}
}

func (se *ScriptExecutor) Bind() {
	for name, fn := range se.BindFuncs {
		se.Runtime.Set(name, fn)
	}

}

func (se *ScriptExecutor) Entry(name string, entry interface{}) error {
	rawentry := se.Runtime.Get(name)
	if rawentry == nil {
		return easyerr.NotFound()
	}

	return se.Runtime.ExportTo(rawentry, entry)
}

func (se *ScriptExecutor) Clear() {
	for name := range se.BindFuncs {
		se.Runtime.Set(name, nil)
	}
}

package hooked

import (
	"github.com/dop251/goja"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/thoas/go-funk"
)

type Hook struct {
	Name      string                 `json:"name,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Targets   []string               `json:"targets,omitempty"`
	Value     string                 `json:"value,omitempty"`
	LifeCycle string                 `json:"lifecycle,omitempty"`
	Options   map[string]interface{} `json:"options,omitempty"`
}

type Hooked struct {
	hooks             []Hook
	runtime           *goja.Runtime
	currentLifcycle   string
	currentTarget     string
	skipLfMethodCheck bool
}

func New(b rtypes.Bindings, rt *goja.Runtime) *Hooked {
	h := &Hooked{
		hooks:             []Hook{},
		runtime:           rt,
		currentLifcycle:   "",
		currentTarget:     "",
		skipLfMethodCheck: false,
	}

	return h
}

type LifecycleMethod struct {
	Name      string
	Lifecyles []string
	Method    func(args ...interface{}) (interface{}, error)
}

func (h *Hooked) AddLifecycleMethod(opts LifecycleMethod) {
	err := h.runtime.Set(opts.Name, func(vals ...interface{}) (interface{}, interface{}) {
		if !h.skipLfMethodCheck || len(opts.Lifecyles) > 0 {
			if !funk.ContainsString(opts.Lifecyles, h.currentLifcycle) {
				return nil, "Not Allowed"
			}
		}

		resp, err := opts.Method(vals...)
		if err != nil {
			return nil, err.Error()
		}
		return resp, nil
	})

	if err != nil {
		panic(err)
	}

}

func (h *Hooked) ApplyHook(lifecycle, target string, val interface{}) error {

	for _, hook := range h.hooks {
		if hook.LifeCycle != lifecycle {
			continue
		}

		if !funk.ContainsString(hook.Targets, target) {
			continue
		}

		h.currentLifcycle = lifecycle
		h.currentTarget = target

		switch hook.Type {
		case "script":
			return h.doGojaHook(val, hook)
		default:
			panic("not implemented")
		}
	}

	return nil
}

func (h *Hooked) doGojaHook(val interface{}, hook Hook) error {
	// fixme => apply actual hook here
	// hook.Value should contain function name

	return nil

}

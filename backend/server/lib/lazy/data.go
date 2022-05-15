package lazydata

import (
	"encoding/json"

	"github.com/dop251/goja"
	"github.com/mitchellh/mapstructure"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"gitlab.com/mr_balloon/golib/hmap"
)

// Data is a container type that holds byets/runtime_types/opaque type
// which will be initilized JIT to apporopate type
type data struct {
	inner interface{}
	ctx   *goja.Runtime
}

func New(payload interface{}, ctx *goja.Runtime) *data {
	return &data{
		inner: payload,
		ctx:   ctx,
	}
}

func (d *data) As(target interface{}) error {
	switch inner := d.inner.(type) {
	case []byte:
		return json.Unmarshal(inner, target)
	default:
		return mapstructure.Decode(d.inner, target)
	}
}

func (d *data) AsBytes() ([]byte, error) {
	if d.ctx != nil {
		var inner interface{}
		err := d.exportRuntimevalue(inner)
		if err != nil {
			return nil, err
		}

		return json.Marshal(inner)
	}

	switch inner := d.inner.(type) {
	case []byte:
		return inner, nil
	default:
		return json.Marshal(d.inner)
	}
}

func (d *data) AsHMap() (hmap.H, error) {
	switch inner := d.inner.(type) {
	case []byte:
		var h hmap.H
		err := json.Unmarshal(inner, &h)
		if err != nil {
			return nil, err
		}
		return h, nil

	case map[string]interface{}:
		return hmap.H(inner), nil
	default:
		return nil, easyerr.NotImpl()
	}

}

func (d *data) exportRuntimevalue(target interface{}) error {
	runtime := d.ctx

	val, ok := d.inner.(goja.Value)
	if !ok {
		return easyerr.NotImpl()
	}

	err := runtime.ExportTo(val, target)
	if err != nil {
		return err
	}
	return nil

}

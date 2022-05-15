package lib

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/dop251/goja"
	"github.com/icza/dyno"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"gitlab.com/mr_balloon/golib/hmap"
)

var (
	ErrInvalidCtx = errors.New("invalid context")
)

type mode uint8

const (
	modeUnknown    mode = 0
	modeGojaUninit mode = 1
	modeByte       mode = 2
	modeHmap       mode = 3
)

type LazyMap struct {
	modeTag mode
	bdata   []byte
	data    hmap.H
	value   interface{}
	context interface{}
}

func New(data interface{}) LazyMap {
	switch ty := data.(type) {
	case []byte:
		return FromBytes(ty)
	case map[string]interface{}:
		return FromAny(ty)
	case map[string][]string:
		dh := make(hmap.H)
		for k, v := range ty {
			dh[k] = v
		}
		return FromAny(dh)

	case map[string]string:
		dh := make(hmap.H)
		for k, v := range ty {
			dh[k] = v
		}
		return FromAny(dh)
	default:
		return FromAny(hmap.H{
			"data": ty,
		})
	}
}

func FromAny(data hmap.H) LazyMap {
	return LazyMap{
		bdata:   nil,
		data:    data,
		modeTag: modeHmap,
	}
}

func FromBytes(bytes []byte) LazyMap {
	return LazyMap{
		bdata:   bytes,
		data:    nil,
		modeTag: modeByte,
	}
}

func FromGojaCtx(val goja.Value, ctx *goja.Runtime) LazyMap {
	return LazyMap{
		modeTag: modeGojaUninit,
		bdata:   nil,
		data:    nil,
		value:   val,
		context: ctx,
	}
}

func (m *LazyMap) transformGojaCtx(target interface{}) error {

	runtime, ok := m.context.(*goja.Runtime)
	if !ok {
		return ErrInvalidCtx
	}

	val, ok := m.value.(goja.Value)
	if !ok {
		return ErrInvalidCtx
	}
	err := runtime.ExportTo(val, target)
	if err != nil {
		return ErrInvalidCtx
	}
	m.modeTag = modeHmap
	return nil
}

func (m *LazyMap) AsBytes() ([]byte, error) {
	if m.modeTag == modeByte {
		return m.bdata, nil
	}
	return json.Marshal(m.data)
}

func (m *LazyMap) AsHMap() (hmap.H, error) {
	if m.modeTag == modeHmap {
		return m.data, nil
	} else if m.modeTag == modeGojaUninit {
		err := m.transformGojaCtx(&m.data)
		if err != nil {
			return nil, err
		}
	}

	var h hmap.H
	err := json.Unmarshal(m.bdata, &h)
	return h, err
}

func (m *LazyMap) AsStruct(i interface{}) error {
	if m.modeTag == modeHmap {
		return m.data.Fill(i)
	} else if m.modeTag == modeGojaUninit {
		return m.transformGojaCtx(i)
	}
	return json.Unmarshal(m.bdata, i)
}

func (m *LazyMap) Set(value interface{}, path ...string) error {
	if m.modeTag == modeHmap {
		return m.setBytes(value, concat(path...))
	} else if m.modeTag == modeGojaUninit {
		err := m.transformGojaCtx(&m.data)
		if err != nil {
			return err
		}
	}
	return dyno.SSet(m.data, value, path...)
}

func (m *LazyMap) Get(path ...string) (interface{}, error) {
	if m.modeTag == modeHmap {
		r := gjson.GetBytes(m.bdata, concat(path...))
		if r.Exists() {
			return nil, easyerr.NotFound()
		}
		return r.Value(), nil
	} else if m.modeTag == modeGojaUninit {
		err := m.transformGojaCtx(&m.data)
		if err != nil {
			return nil, err
		}
	}

	return dyno.SGet(m.data, path...)
}

func (m *LazyMap) setBytes(value interface{}, path string) error {
	nb, err := sjson.SetBytes(m.bdata, path, value)
	if err != nil {
		return err
	}
	m.bdata = nb
	return nil
}

func (m *LazyMap) Delete(path ...string) error {
	if m.modeTag == modeHmap {
		return m.deleteBytes(concat(path...))
	} else if m.modeTag == modeGojaUninit {
		err := m.transformGojaCtx(&m.data)
		if err != nil {
			return err
		}
	}

	ipaths := make([]interface{}, 0, len(path))
	for i, p := range path {
		if i == 0 {
			continue
		}

		ipaths = append(ipaths, p)
	}
	return dyno.Delete(m.data, path[0], ipaths...)
}

func (m *LazyMap) deleteBytes(path string) error {
	nb, err := sjson.DeleteBytes(m.bdata, path)
	if err != nil {
		return err
	}
	m.bdata = nb
	return nil
}

func (m *LazyMap) Booli(key string, val ...bool) bool {
	if m.modeTag == modeHmap {
		r := gjson.GetBytes(m.bdata, key)
		if r.Exists() {
			return r.Bool()
		}
		if len(val) > 0 {
			return val[0]
		}
		return false
	} else if m.modeTag == modeGojaUninit {
		err := m.transformGojaCtx(&m.data)
		if err != nil {
			return false
		}
	}
	return m.data.Booli(key, val...)
}

func (m *LazyMap) Floati(key string, val ...float64) float64 {
	if m.modeTag == modeHmap {
		r := gjson.GetBytes(m.bdata, key)
		if r.Exists() {
			return r.Float()
		}
		if len(val) > 0 {
			return val[0]
		}
		return 0
	} else if m.modeTag == modeGojaUninit {
		err := m.transformGojaCtx(&m.data)
		if err != nil {
			return 0
		}
	}
	return m.data.Float64i(key, val...)
}

func (m *LazyMap) Int64i(key string, val ...int64) int64 {
	if m.modeTag == modeHmap {
		r := gjson.GetBytes(m.bdata, key)
		if r.Exists() {
			return r.Int()
		}
		if len(val) > 0 {
			return val[0]
		}
		return 0
	} else if m.modeTag == modeGojaUninit {
		err := m.transformGojaCtx(&m.data)
		if err != nil {
			return 0
		}

	}

	return m.data.Int64i(key, val...)
}

func (m *LazyMap) Stringi(key string, val ...string) string {
	if m.modeTag == modeHmap {
		r := gjson.GetBytes(m.bdata, key)
		if r.Exists() {
			return r.String()
		}
		if len(val) > 0 {
			return val[0]
		}
		return ""
	} else if m.modeTag == modeGojaUninit {
		err := m.transformGojaCtx(&m.data)
		if err != nil {
			return ""
		}

	}
	return m.data.Stringi(key, val...)

}
func (m *LazyMap) Uint64i(key string, val ...uint64) uint64 {
	if m.modeTag == modeHmap {
		r := gjson.GetBytes(m.bdata, key)
		if r.Exists() {
			return r.Uint()
		}
		if len(val) > 0 {
			return val[0]
		}
		return 0
	} else if m.modeTag == modeGojaUninit {
		err := m.transformGojaCtx(&m.data)
		if err != nil {
			return 0
		}

	}

	return m.data.Uint64i(key, val...)
}

func concat(ss ...string) string {
	if len(ss) == 1 {
		return ss[0]
	}

	var buf strings.Builder

	for _, s := range ss {
		buf.WriteString(s)
	}
	return buf.String()
}

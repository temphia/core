package dummy

import (
	"encoding/json"
	"time"

	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/event"
)

var _ rtypes.Executor = (*Dummy)(nil)

func New(opts rtypes.ExecutorOption) (rtypes.Executor, error) {
	return &Dummy{
		Binder:   opts.Binder,
		PlugId:   opts.PlugId,
		AgentId:  opts.AgentId,
		Slug:     opts.Slug,
		ExecType: opts.ExecType,
	}, nil
}

type Dummy struct {
	Binder   rtypes.Bindings
	PlugId   string
	AgentId  string
	Slug     string
	ExecType string
}

type PingResp struct {
	Action  string          `json:"action,omitempty"`
	Time    time.Time       `json:"time,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (d *Dummy) Process(ev *event.Request) (*event.Response, error) {
	var out []byte
	switch ev.Name {
	case "ping":
		out1, err := d.actionPing(ev)
		if err != nil {
			return nil, err
		}
		out = out1
	default:
		return nil, easyerr.Error("action not found")
	}

	return &event.Response{
		Vars:    map[string]interface{}{},
		Payload: out,
	}, nil
}

func (d *Dummy) actionPing(ev *event.Request) ([]byte, error) {
	ping := PingResp{
		Action:  ev.Name,
		Time:    time.Now(),
		Message: "hello",
		Data:    ev.Data,
	}

	return json.Marshal(&ping)
}

func (d *Dummy) actionAddState(ev *event.Request) ([]byte, error) {
	pkv := d.Binder.GetPlugKVBindings()
	err := pkv.Set(0, "key1", "value", nil)
	if err != nil {
		return nil, err
	}

	return []byte(`state added`), nil
}

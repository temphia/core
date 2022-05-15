package dynhub

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/job"
)

/*
	hooks types
		- on_row_ctx
		- on_table_ctx
		- on_before_mod
		- on_after_mod
*/

type DataEventReq struct {
	Id    int64                  `json:"id,omitempty"`
	Type  string                 `json:"type,omitempty"`
	Group string                 `json:"group,omitempty"`
	Table string                 `json:"table,omitempty"`
	Data  map[string]interface{} `json:"data,omitempty"`
}

func (d *dynSource) OnBeforeMod(tenant, plug, agent, handler string, event DataEventReq) (map[string]interface{}, error) {

	out, err := json.Marshal(&event)
	if err != nil {
		return nil, err
	}

	j := &job.Job{
		PlugId:            plug,
		AgentId:           agent,
		Namespace:         tenant,
		EventId:           xid.New().String(),
		EventType:         "data_event",
		EventName:         handler,
		RequestVars:       map[string]interface{}{},
		Claim:             nil,
		Payload:           out,
		PendingPrePolicy:  true,
		PendingPostPolicy: true,
		Plug:              nil,
		Agent:             nil,
		Resources:         nil,
		Loaded:            false,
	}

	runtime := d.hub.engine.GetRuntime()

	j.Add()
	runtime.Shedule(j)
	j.Wait()

	resp, err := j.Result()
	if err != nil {
		return nil, err
	}

	pp.Println(resp)

	return nil, nil
}

func (d *dynSource) OnAfterMod(event DataEventReq) error {
	return nil
}

package binder

import (
	"encoding/json"
	"time"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
)

func (b *Binder) Log(msg string) {
	b.logger.Debug().Msg(msg)

	b.logDebugRoom(&DebugMessage{
		messages: []string{msg},
		Xid:      xid.New().String(),
		PlugId:   b.plugId,
		AgentId:  b.agentId,
	})
}

func (b *Binder) LazyLog(msgs []string) {
	b.logger.Debug().Strs("batched_log", msgs).Send()

	b.logDebugRoom(&DebugMessage{
		messages: msgs,
		Xid:      xid.New().String(),
		PlugId:   b.plugId,
		AgentId:  b.agentId,
	})
}

func (b *Binder) Sleep(msec int) {
	time.Sleep(time.Millisecond * time.Duration(msec))
}

func (b *Binder) GetSelfFile(file string) ([]byte, error) {
	return b.factory.Pacman.BprintGetBlob(b.namespace, b.job.Plug.BprintId, file)
}

type DebugMessage struct {
	Xid      string
	PlugId   string
	AgentId  string
	messages []string
}

func (b *Binder) logDebugRoom(msg *DebugMessage) {
	out, err := json.Marshal(msg)
	if err != nil {
		pp.Println(err)
		return
	}

	b.factory.Sockd.SendBroadcast(b.namespace, "plugs_dev", []string{}, out)
}

func (b *Binder) GetApp() interface{} {
	return b.factory.App
}

func (b *Binder) GetRequestData() interface{} {
	return b.job.Payload
}

func (b *Binder) GetRequestVar(vname string) interface{} {
	return b.job.RequestVars[vname]
}

func (b *Binder) GetRequestVars() map[string]interface{} {
	return b.job.RequestVars
}

func (b *Binder) SetResponse(resp interface{}) {
	b.resp = resp
}

func (b *Binder) SetResponseVar(name string, value interface{}) {
	if b.respVars == nil {
		b.respVars = make(map[string]interface{})
	}

	b.respVars[name] = value
}

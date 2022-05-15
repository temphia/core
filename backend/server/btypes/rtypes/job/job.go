package job

import (
	"sync"

	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
)

type Job struct {
	PlugId    string
	AgentId   string
	EventId   string
	EventType string
	EventName string
	Namespace string
	// ExecNodeTag       string
	// OriginNode        string
	RequestVars map[string]interface{}
	Claim       interface{}
	Payload     []byte
	// Payload           btypes.Data

	PendingPrePolicy  bool
	PendingPostPolicy bool

	// lazy loadable
	Plug      *entities.Plug
	Agent     *entities.Agent
	Resources map[string]*entities.Resource
	Loaded    bool

	// result stuff
	result    *event.Response
	resultErr error

	wg sync.WaitGroup
}

func (j *Job) Err(err error) {
	j.resultErr = err
	j.wg.Done()
}

func (j *Job) Ok(resp *event.Response) {
	j.result = resp

	pp.Println(j.wg)

	j.wg.Done()
}

func (j *Job) Add() {
	j.wg.Add(1)
}

func (j *Job) Wait() {

	pp.Println(j.wg)

	j.wg.Wait()
}

func (j *Job) Result() (*event.Response, error) {
	return j.result, j.resultErr
}

func (j *Job) AsEvent() *event.Request {
	return &event.Request{
		Id:   j.EventId,
		Type: j.EventType,
		Name: j.EventName,
		Vars: j.RequestVars,
		Data: j.Payload,
	}

}

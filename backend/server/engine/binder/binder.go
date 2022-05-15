package binder

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/event"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/job"
	"gitlab.com/mr_balloon/golib"
)

var (
	_ rtypes.Bindings = (*Binder)(nil)
)

type Binder struct {
	factory   Factory
	namespace string
	plugId    string
	executor  rtypes.Executor

	// These should be set @Execute
	agentId   string
	eventId   string
	resources map[string]*entities.Resource
	ctx       context.Context // fixme

	ipcBinder    ipcBinder
	plugKvBinder plugKvBinder
	logger       zerolog.Logger
	job          *job.Job

	resp     interface{}
	respVars map[string]interface{}
}

func (b *Binder) SetExec(exec rtypes.Executor) {
	b.executor = exec
}

var NoPanicWrap = true

func (b *Binder) Execute() {
	defer b.clearPrivate()

	if err := b.precheck(); err != nil {
		b.logger.Error().Err(err).Msg("Precheck failed")
		b.job.Err(err)
		return
	}

	b.logger.Info().Msg("Before processing event")

	var eresp *event.Response
	var err error

	if NoPanicWrap {
		eresp, err = b.executor.Process(b.job.AsEvent())
	} else {
		perr := golib.PanicWrapper(func() {
			eresp, err = b.executor.Process(b.job.AsEvent())
		})

		if perr != nil {
			b.logger.Error().Err(err).Msg("Executor Panicked while processing event")
			b.job.Err(perr)
			return
		}
	}

	if err != nil {
		b.logger.Error().Err(err).Msg("Executor threw a error")
		b.job.Err(err)
		return
	}

	b.logger.Info().Msg("Executor processed event sucessfully")

	if eresp.Payload == nil && b.resp != nil {
		eresp.Payload = b.resp
	}

	if b.respVars != nil {
		if eresp.Vars == nil {
			eresp.Vars = b.respVars
		} else {
			for k, v := range b.respVars {
				_, ok := eresp.Vars[k]
				if ok {
					continue
				}
				eresp.Vars[k] = v
			}
		}
	}

	b.job.Ok(eresp)
}

func (b *Binder) AttachJob(j *job.Job) {
	b.job = j

	if j.AgentId != b.agentId {
		b.clear()
		b.resources = j.Resources
		b.ctx = context.Background()
		b.agentId = j.AgentId
	}
	b.logger = b.factory.Logger.
		With().
		Str("tenant_id", b.namespace).
		Str("plug_id", b.namespace).
		Str("agent_id", b.agentId).
		Str("event_id", j.EventId).Logger()

	b.eventId = j.EventId
}

// private
func (b *Binder) clear() {
	b.ctx = nil
	b.resources = nil
	b.agentId = ""
	b.eventId = ""

	b.ipcBinder.clear()
	// b.stateBinder.clearTxns()
}

func (b *Binder) clearPrivate() {
	b.job = nil
	// b.stateBinder.clearTxns()
}

func (b *Binder) precheck() error {
	if b.executor == nil {
		return easyerr.Error("executor is empty")
	}

	if b.resources == nil {
		return easyerr.Error("resources not loaded/ empty")

	}

	return nil
}

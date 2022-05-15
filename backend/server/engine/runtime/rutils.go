package runtime

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/rtypes/job"
	"github.com/temphia/core/backend/server/engine/binder"
	"github.com/temphia/core/backend/server/helpers/plugloader"
	"github.com/ztrue/tracerr"
)

func (r *runtime) getBinder(j *job.Job) (rtypes.ExecutorBinder, error) {
	if !j.Loaded {
		err := r.loadJob(j)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}
	}

	excbinder := r.pool.borrow(j.Namespace, j.PlugId, j.AgentId)
	if excbinder != nil {

		excbinder.AttachJob(j)

		return excbinder, nil
	}

	eb, ok := r.execBuilders[j.Plug.Executor]
	if !ok {
		fmt.Println(r.execBuilders)
		return nil, easyerr.Error("Executor builder not found")
	}

	bind := r.binderFactory.New(binder.BinderOptions{
		Namespace: j.Namespace,
		PlugId:    j.PlugId,
	})

	bind.AttachJob(j)

	exec, err := eb.Instance(rtypes.ExecutorOption{
		Binder:   bind,
		PlugId:   j.PlugId,
		AgentId:  j.AgentId,
		Slug:     j.Plug.Id,
		ExecType: j.Plug.Executor,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	bind.SetExec(exec)
	return bind, nil
}

func (r *runtime) setBinder(j *job.Job, bind rtypes.ExecutorBinder) {
	r.pool.returnn(j.Namespace, j.PlugId, j.AgentId, bind)
}

func (r *runtime) loadJob(j *job.Job) error {
	pp.Println("@=>", j.Namespace, j.PlugId, j.AgentId)

	data, err := plugloader.Load(r.syncer, j.Namespace, j.PlugId, j.AgentId)

	if err != nil {
		return tracerr.Wrap(err)
	}
	j.Agent = data.Agent
	j.Plug = data.Plug
	j.Resources = data.Resources
	j.Loaded = true
	return nil
}

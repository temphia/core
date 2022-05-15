package engine

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/ztrue/tracerr"

	"github.com/temphia/temphia/backend/server/btypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/job"
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/temphia/temphia/backend/server/btypes/store"
	"github.com/temphia/temphia/backend/server/engine/runtime"
	"github.com/temphia/temphia/backend/server/lib/apiutils"
	"github.com/temphia/temphia/backend/server/registry"
)

type Engine struct {
	app        btypes.App
	runtime    rtypes.Runtime
	signer     service.Signer
	syncer     store.SyncDB
	AssetStore store.AssetStore
	pacman     service.Pacman
	logger     zerolog.Logger
	builders   map[string]rtypes.ExecutorBuilder
}

func New(_app btypes.App, logger zerolog.Logger) *Engine {

	return &Engine{
		app:        _app,
		runtime:    nil,
		signer:     nil,
		syncer:     nil,
		AssetStore: _app.Data(),
		pacman:     nil,
		logger:     logger,
	}

}

func (e *Engine) Run() error {

	e.runtime = runtime.New(e.app, e.logger)
	e.signer = e.app.Signer().(service.Signer)
	e.syncer = e.app.CoreHub()
	e.pacman = e.app.Pacman().(service.Pacman)

	reg := e.app.Registry().(*registry.Registry)
	e.builders = reg.GetExecutors()

	return e.runtime.Run()
}

func (e *Engine) GetRuntime() rtypes.Runtime {
	return e.runtime
}

func (e *Engine) OnConsoleExec(tenantId, plugId, agentId, action string, ctx *gin.Context) {

	claim, err := e.tryExtractClaim(tenantId, ctx)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	out, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	j := &job.Job{
		PlugId:    plugId,
		AgentId:   agentId,
		Namespace: tenantId,
		EventId:   xid.New().String(),
		EventType: "console",
		EventName: action,
		RequestVars: map[string]interface{}{
			"http_headers": ctx.Request.Header,
			"http_url":     ctx.Request.URL.String(),
			"http_method":  http.MethodPost,
		},
		Claim:             claim,
		Payload:           out,
		PendingPrePolicy:  true,
		PendingPostPolicy: true,
		Plug:              nil,
		Agent:             nil,
		Resources:         nil,
		Loaded:            false,
	}

	pp.Println(j)

	j.Add()
	e.runtime.Shedule(j)
	j.Wait()

	eresp, err := j.Result()
	if err != nil {
		tracerr.PrintSourceColor(err, 10)
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	pp.Println("Writing response =>")

	switch data := eresp.Payload.(type) {
	case []byte:
		pp.Println(string(data))
		ctx.Header("Content-Type", "application/json")
		apiutils.WriteBinary(ctx, data)
		return
	default:

		pp.Println((data))
		ctx.Header("Content-Type", "application/json")
		apiutils.WriteJSON(ctx, data, nil)
		return
	}
}

func (e *Engine) OnRawExec(tenantId, plugId, agentId, action string, ctx *gin.Context) {

	j := &job.Job{
		PlugId:    plugId,
		AgentId:   agentId,
		Namespace: tenantId,
		EventId:   xid.New().String(),
		EventType: "console",
		EventName: action,
		RequestVars: map[string]interface{}{
			"http_headers": ctx.Request.Header,
			"http_url":     ctx.Request.URL.String(),
			"http_method":  ctx.Request.Method,
		},
		Claim:             nil,
		Payload:           nil,
		PendingPrePolicy:  true,
		PendingPostPolicy: true,
		Plug:              nil,
		Agent:             nil,
		Resources:         nil,
		Loaded:            false,
	}
	e.runtime.Shedule(j)
	j.Wait()

	eresp, err := j.Result()
	if err != nil {
		return
	}

	// fixme => write resp
	pp.Println(eresp)

}

func (e *Engine) OnFedExec(tenantId, resourceId, action string, ctx *gin.Context) {

	return
}

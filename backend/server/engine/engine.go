package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/temphia/core/backend/server/engine/runtime"
	"github.com/temphia/core/backend/server/registry"
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

func (e *Engine) ServerLaunchExec(tenantId, plugId, agentId, mode string, arg interface{}, resp interface{}) error {
	return e.serverLaunchExec(tenantId, plugId, agentId, mode, arg, resp)
}

func (e *Engine) ClientLaunchExec(tenantId, plugId, agentId, mode string, ctx *gin.Context) {
	e.clientLaunchExec(tenantId, plugId, agentId, mode, ctx)
}

func (e *Engine) ExecAction(tenantId, plugId, agentId, action string, ctx *gin.Context) {
	e.plugAction(tenantId, plugId, agentId, action, ctx)
}

func (e *Engine) ServePlugFile(tenantId, plugId, agentId, file string, ctx *gin.Context) {
	e.servePlugFile(tenantId, plugId, agentId, file, ctx)
}

func (e *Engine) ServeExecutorFile(tenantId, plugId, agentId, loader string, ctx *gin.Context) {
	e.serveExecutorFile(tenantId, plugId, agentId, loader, ctx)
}

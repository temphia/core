package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/lib/apiutils"
	"github.com/temphia/core/backend/server/services/sockcore/transports"
)

func (r *R) EngineExecConsole(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	action := ctx.Param("action")
	r.engine.OnConsoleExec(tenantId, plugId, agentId, action, ctx)
}

func (r *R) EngineExecRaw(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	action := ctx.Param("action")
	r.engine.OnRawExec(tenantId, plugId, agentId, action, ctx)
}

// fixme => remove this
func (r *R) EngineExecFed(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	resourceId := ctx.Param("resource_id")
	action := ctx.Param("action")
	r.engine.OnFedExec(tenantId, resourceId, action, ctx)
}

func (r *R) EngineSRLauncher(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	r.engine.LaunchSubOrigin(tenantId, plugId, agentId, ctx)
}

func (r *R) EngineIFrameLauncher(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")

	// fixme => also extract other ctx

	r.engine.LaunchIFrame(tenantId, plugId, agentId, ctx)
}

func (r *R) EngineServe(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	file := ctx.Param("file")
	r.engine.Serve(tenantId, plugId, agentId, file, ctx)
}

func (r *R) EngineExecLoaderScript(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	loader := ctx.Param("loader")

	out, err := r.engine.ExecutorFile(tenantId, plugId, agentId, fmt.Sprintf("%s_loader.js", loader))
	if err != nil {
		pp.Println("@=>", err)
		return
	}
	ctx.Writer.Header().Set("Content-Type", "application/javascript")
	ctx.Writer.Write(out)
}

func (r *R) EngineExecLoaderStyle(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	loader := ctx.Param("loader")

	out, err := r.engine.ExecutorFile(tenantId, plugId, agentId, fmt.Sprintf("%s_loader.css", loader))
	if err != nil {
		pp.Println("@=>", err)
		return
	}

	ctx.Writer.Header().Set("Content-Type", "text/css")
	ctx.Writer.Write(out)
}

func (r *R) PlugSocket(c *gin.Context) {
	tenant := c.Param("tenant_id")
	agent := c.Param("agent_id")
	plug := c.Param("plug_id")

	// token := c.Query("token")

	pp.Println(agent, plug)

	conn, err := transports.NewConnWS(c, xid.New().String())
	if err != nil {
		apiutils.WriteErr(c, err.Error())
		return
	}

	connTags := []string{
		fmt.Sprintf("plug_%s", plug),
	}

	err = r.sockd.NewConnection(&service.ConnOptions{
		NameSpace: tenant,
		Conn:      conn,
		Expiry:    10000,
		PreJoinRooms: map[string][]string{
			btypes.ROOM_PLUG_DEV: connTags,
		},
	})
	if err != nil {
		apiutils.WriteErr(c, err.Error())
		return
	}
}

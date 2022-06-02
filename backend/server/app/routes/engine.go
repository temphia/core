package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/core/backend/server/lib/apiutils"
	"github.com/temphia/core/backend/server/services/sockcore/transports"
	"github.com/temphia/core/backend/server/services/sockdhub"
)

func (r *R) EngineExecConsole(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	action := ctx.Param("action")
	r.engine.ExecAction(tenantId, plugId, agentId, action, ctx)
}

func (r *R) EngineLaunchExecHTML(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")

	// fixme => also extract other ctx

	r.engine.ClientLaunchExec(tenantId, plugId, agentId, "ssr", ctx)
}

func (r *R) EngineLaunchExec(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")

	// fixme => also extract other ctx

	r.engine.ClientLaunchExec(tenantId, plugId, agentId, "", ctx)
}

func (r *R) EngineServe(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	file := ctx.Param("file")
	pp.Println(tenantId, plugId, agentId, file)
	//	r.engine.ServePlugFile(tenantId, plugId, agentId, file, ctx)
}

func (r *R) EngineExecLoaderScript(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	loader := ctx.Param("loader")

	r.engine.ServeExecutorFile(tenantId, plugId, agentId, fmt.Sprintf("%s_loader.js", loader), ctx)
}

func (r *R) EngineExecLoaderStyle(ctx *gin.Context) {
	tenantId := ctx.Param("tenant_id")
	plugId := ctx.Param("plug_id")
	agentId := ctx.Param("agent_id")
	loader := ctx.Param("loader")

	r.engine.ServeExecutorFile(tenantId, plugId, agentId, fmt.Sprintf("%s_loader.css", loader), ctx)
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

	r.sockdhub.AddPlugConnection(sockdhub.PlugConnOptions{
		TenantId: tenant,
		UserId:   "",
		GroupId:  "",
		DeviceId: "",
		Plug:     plug,
		Conn:     conn,
	})

	if err != nil {
		apiutils.WriteErr(c, err.Error())
		return
	}
}

package engine

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/core/backend/server/btypes/models/vmodels"
	"github.com/temphia/core/backend/server/btypes/rtypes/job"
	"github.com/temphia/core/backend/server/lib/apiutils"
	"github.com/ztrue/tracerr"
)

func (e *Engine) clientLaunchExec(tenantId, plugId, agentId, mode string, ctx *gin.Context) {
	if mode == "ssr" {
		e.clientLaunchExecSSR(tenantId, plugId, agentId, ctx)
		return
	}

	pp.Println("@@@@=>", tenantId, plugId, agentId)

	agent, err := e.syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	// plug, err := e.syncer.PlugGet(tenantId, plugId)
	// if err != nil {
	// 	apiutils.WriteErr(ctx, err.Error())
	// 	return
	// }

	resp := &vmodels.LoaderOptions{
		BaseURL:      baseURL(ctx),
		Token:        "",
		EntryName:    agent.EntryName,
		ExecLoader:   agent.ExecLoader,
		JSPlugScript: agent.EntryScript,
		Plug:         plugId,
		Agent:        agentId,
		StyleFile:    agent.EntryStyle,
		ExtScripts:   nil,
	}

	apiutils.WriteJSON(ctx, resp, nil)

}

func (e *Engine) plugAction(tenantId, plugId, agentId, action string, ctx *gin.Context) {
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

func (e *Engine) servePlugFile(tenantId, plugId, agentId, file string, ctx *gin.Context) {
	plug, err := e.syncer.PlugGet(tenantId, plugId)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	agent, err := e.syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}
	actualFile := agent.ServeFiles[file]

	if actualFile == "" {
		pp.Println(tenantId, plugId, agentId)
		return
	}

	out, err := e.pacman.BprintGetBlob(tenantId, plug.BprintId, actualFile)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	e.writeFile(file, out, ctx)
}

func (e *Engine) serveExecutorFile(tenantId, plugId, agentId, file string, ctx *gin.Context) {
	plug, err := e.syncer.PlugGet(tenantId, plugId)
	if err != nil {
		return
	}

	builder, ok := e.builders[plug.Executor]
	if !ok {
		return
	}

	out, err := builder.ExecFile(file)
	if err != nil {
		return
	}

	e.writeFile(file, out, ctx)
}

func (e *Engine) writeFile(file string, data []byte, ctx *gin.Context) {

	ffiles := strings.Split(file, ".")

	switch ffiles[1] {
	case "js":
		ctx.Writer.Header().Set("Content-Type", "application/javascript")
	default:
		ctx.Writer.Header().Set("Content-Type", "text/css")
	}
	ctx.Writer.Write(data)
}

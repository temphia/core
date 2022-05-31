package engine

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/models/vmodels"
	"github.com/temphia/core/backend/server/lib/apiutils"
)

// fixme => x-content-security-policy: frame-ancestors 'self' https://mycourses.w3schools.com;
// Referer: https://example/launcher/<ticket>

// suborigin launcher
func (e *Engine) clientLaunchExecSSR(tenantId, plugId, agentId string, ctx *gin.Context) {
	token := refererToken(ctx)
	pp.Println(token)

	loader, err := e.AssetStore.GetTemplate("suborigin_loader.js")
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	agent, err := e.syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	plug, err := e.syncer.PlugGet(tenantId, plugId)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	// fixme => temporary hack

	if plug.Executor == "simplewizard" {
		agent.ExecLoader = "simplewizard.main"
	}

	rData, err := buildSubOriginTemplate(&vmodels.SubOriginData{
		LoaderJS:       string(loader),
		LoaderOptsJSON: "",
		BaseURL:        baseURL(ctx),
		Token:          "",
		Plug:           plugId,
		Agent:          agentId,
		EntryName:      agent.EntryName,
		ExecLoader:     agent.ExecLoader,
		JSPlugScript:   agent.EntryScript,
		StyleFile:      agent.EntryStyle,
		ExtScripts:     nil,
	})
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	ctx.Writer.Write(rData)
}

func refererToken(ctx *gin.Context) string {
	referer := ctx.Request.Header.Get("Referer")
	url, err := url.Parse(referer)
	if err == nil && url != nil {
		ticket := url.Query().Get("referer_token")
		if ticket != "" {
			return ticket
		}
	}
	return ctx.Query("referer_token")
}

func baseURL(ctx *gin.Context) string {
	return fmt.Sprintf("http://%s/z/api/%s/v1/", ctx.Request.Host, ctx.Param("tenant_id"))
}

func buildSubOriginTemplate(renderOpts *vmodels.SubOriginData) ([]byte, error) {

	err := renderOpts.BuildJSONOpts()
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	buf.Write([]byte(`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
		<script async="true"> window["__loader_options__"] =
	`),
	)
	buf.WriteString(renderOpts.LoaderOptsJSON)
	buf.Write([]byte(`</script> <script async="true">`))
	buf.WriteString(renderOpts.LoaderJS)
	buf.Write([]byte(`</script>`))

	if renderOpts.ExecLoader != "" {
		buf.WriteString(
			fmt.Sprintf(
				`<script src="%sengine/%s/%s/executor/%s/loader.js"></script>`,
				renderOpts.BaseURL,
				renderOpts.Plug,
				renderOpts.Agent,
				renderOpts.ExecLoader,
			),
		)
	}

	buf.WriteString(
		fmt.Sprintf(
			`<script src="%sengine/%s/%s/serve/%s"></script>`,
			renderOpts.BaseURL,
			renderOpts.Plug,
			renderOpts.Agent,
			renderOpts.JSPlugScript,
		),
	)

	buf.Write([]byte(`
	</head>
	<body>
	<div id="plugroot" style="height:100vh;"></div>
	</body>
	</html>	
	`))

	return buf.Bytes(), nil
}

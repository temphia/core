package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/rs/zerolog"
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/models/vmodels"
	"github.com/temphia/core/backend/server/lib/apiutils"
	"github.com/temphia/core/backend/server/lib/apiutils/request"
	"github.com/ztrue/tracerr"
)

func (r *R) AuthIndex2(ctx *gin.Context) {
	var buf bytes.Buffer

	tenantId := extractTenant(ctx)
	if tenantId == "" {
		apiutils.WriteErr(ctx, "Could not determine tenant")
		return
	}

	stoken, err := r.siteToken(tenantId, ctx)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	data := &vmodels.SiteData{
		SiteToken: stoken,
		ApiURL:    fmt.Sprintf("http://%s/api/%s/v1", ctx.Request.Host, tenantId),
		TenantId:  tenantId,
	}

	tdata, err := json.Marshal(data)
	if err != nil {
		apiutils.WriteErr(ctx, err.Error())
		return
	}

	ctx.SetCookie("tenant_id", data.TenantId, 60*60*24*30, "/", "", true, true)
	// also attach device_id

	buf.Write([]byte(`<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="utf-8" />
			<meta name="viewport" content="width=device-width,initial-scale=1" />
			<title>Admin Console</title>
			<script>
			window.__temphia_site_data__ = `))
	buf.Write(tdata)
	buf.Write([]byte(`</script>
		<link rel="icon" type="image/png" href="/favicon.png" />
		<link rel="stylesheet"  type="text/css" href="/assets/page_auth.css" />
		<script defer src="/assets/page_auth.js"></script>
		</head>
		<body></body>
		</html>`))

	ctx.Writer.Write(buf.Bytes())
}

func (r *R) Login(c *gin.Context) {

	pp.Println("@@@LOGIN")

	req := vmodels.LoginRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		apiutils.WriteErr(c, err.Error())
		return
	}

	siteClaim := &claim.Site{}
	err = r.signer.Parse(req.TenantId, req.SiteToken, siteClaim)
	if err != nil {
		apiutils.WriteErr(c, err.Error())
		return
	}

	if req.TenantId != siteClaim.TenentID {
		apiutils.WriteErr(c, "invalid token")
		return
	}
	r.cAuth.LogInUser(siteClaim.TenentID, req.UserIdendity, req.Password, c)
}

func (r *R) SignUp(c *gin.Context) {

}

// this should give engine service api
// pinned to specific plug/agent
func (r *R) RefreshFromPairToken(c *gin.Context) {

}

func (r *R) RefreshServiceToken(c *gin.Context) {
	opts := vmodels.RefreshReq{}
	err := c.BindJSON(&opts)
	if err != nil {
		r.WriteErr(c, err.Error())
		return
	}

	uclaim := claim.User{}
	err = r.signer.Parse(c.Param("tenant_id"), opts.UserToken, &uclaim)
	if err != nil {
		r.WriteErr(c, err.Error())
		return
	}
	resp := r.cAuth.RefreshService(&uclaim, &opts)
	apiutils.WriteJSON(c, resp, nil)
}

func (r *R) Authed(fn func(request.Ctx)) func(*gin.Context) {

	logger := zerolog.Nop() // fixme

	return func(c *gin.Context) {

		// time.Sleep(time.Duration(rand.Int()%5) * time.Second)

		claim, err := r.parseSessionClaim(c)
		if err != nil {
			tracerr.PrintSourceColor(err)
			return
		}

		c.Header("X-Clacks-Overhead", "Aaron Swartz")

		funcname := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
		logger.Info().Str("method", funcname).Send()

		fn(request.Ctx{
			Session: claim,
			GinCtx:  c,
		})
	}
}

func (r *R) JSONAuthed(fn func(request.Ctx) (interface{}, error)) func(*gin.Context) {
	logger := zerolog.Nop() // fixme

	return func(c *gin.Context) {
		claim, err := r.parseSessionClaim(c)
		if err != nil {
			frames := tracerr.StackTrace(err)
			logger.Debug().Interface("frames", frames).Err(err).Msg("err response")
			apiutils.WriteErr(c, err.Error())
			return
		}

		funcname := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
		logger.Debug().
			Str("method", funcname).
			Str("tenant_id", claim.TenentID).Send()

		resp, err := fn(request.Ctx{
			Session: claim,
			GinCtx:  c,
		})

		if err != nil {
			tracerr.PrintSourceColor(err)
		}

		apiutils.WriteJSON(c, resp, err)

	}
}

func (r *R) Unsafe(fn func(request.Ctx)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		fn(request.Ctx{
			Session: &claim.Session{
				TenentID:  "default0",
				UserID:    "superuser",
				UserGroup: "super_admin",
				Type:      "",
			},
			GinCtx: ctx,
		})
	}

}

func (r *R) parseSessionClaim(c *gin.Context) (*claim.Session, error) {
	tenantId := c.Param("tenant_id")
	sessToken := c.Request.Header.Get("Authorization")
	sessClaim := &claim.Session{}
	err := r.signer.Parse(tenantId, sessToken, sessClaim)
	if err != nil {
		return nil, err
	}

	if sessClaim.Type != claim.CTypeSession {
		panic("wrong type")
	}

	sessClaim.TenentID = tenantId

	return sessClaim, err
}

func extractTenant(ctx *gin.Context) string {

	tenantId := ctx.Query("tenant_id")
	if tenantId != "" {
		return tenantId
	}
	cookie, _ := ctx.Cookie("tenant_id")
	return cookie
}

func (r *R) siteToken(tenantId string, ctx *gin.Context) (string, error) {
	// fixme => check if host is valid

	siteClaim := claim.NewSiteClaim(tenantId, ctx.Request.Host)
	return r.signer.Sign(tenantId, siteClaim)
}

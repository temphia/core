package routes

import (
	"github.com/temphia/core/backend/server/controller/basic"
	"github.com/temphia/core/backend/server/lib/apiutils/request"
)

func (c *R) RepoSources(ctx request.Ctx) {
	resp, err := c.cBasic.ListRepoSources(ctx.Session)
	c.WriteJSON(ctx.GinCtx, resp, err)
}

func (c *R) RepoList(ctx request.Ctx) {

	resp, err := c.cBasic.RepoList(ctx.Session, basic.RepoListReq{
		Source: ctx.GinCtx.Param("repo"),
		Group:  "",
		Tags:   []string{},
	})
	c.WriteJSON(ctx.GinCtx, resp, err)
}

func (c *R) RepoGet(ctx request.Ctx) {
	resp, err := c.cBasic.RepoGet(ctx.Session, basic.RepoGetReq{
		Source: ctx.GinCtx.Param("repo"),
		Group:  ctx.GinCtx.Param("group_id"),
		Slug:   ctx.GinCtx.Param("slug"),
	})
	c.WriteJSON(ctx.GinCtx, resp, err)
}

func (c *R) RepoGetFile(ctx request.Ctx) {
	resp, err := c.cBasic.RepoGetBlob(ctx.Session, basic.RepoGetBlobReq{
		Source: ctx.GinCtx.Param("repo"),
		Group:  ctx.GinCtx.Param("group_id"),
		Slug:   ctx.GinCtx.Param("slug"),
		File:   ctx.GinCtx.Param("file"),
	})
	c.WriteJSON(ctx.GinCtx, resp, err)
}

/*

func (a *RestAPI) PublicList(c *gin.Context) {
	if a.publicSource == nil {
		apiutils.WriteErr(c, "Not Authorized")
		return
	}

	resp, err := a.publicSource.List("", c.Query("group_id"), c.QueryArray("tags")...)
	apiutils.WriteJSON(c, resp, err)
}

func (a *RestAPI) PublicGet(c *gin.Context) {
	if a.publicSource == nil {
		apiutils.WriteErr(c, "Not Authorized")
		return
	}

	resp, err := a.publicSource.Get("", c.Param("slug"), c.Param("group_id"))
	apiutils.WriteJSON(c, resp, err)
}

func (a *RestAPI) PublicGetBlob(c *gin.Context) {
	if a.publicSource == nil {
		apiutils.WriteErr(c, "Not Authorized")
		return
	}

	resp, err := a.publicSource.GetBlob("", c.Param("slug"), c.Param("group_id"), c.Param("file"))
	apiutils.WriteJSON(c, resp, err)
}


*/

package routes

import (
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/models/vmodels"
	"github.com/temphia/core/backend/server/lib/apiutils/request"
)

func (r *R) ResourceCreate(ctx request.Ctx) {
	res := &entities.Resource{}
	err := ctx.GinCtx.BindJSON(res)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	res.TenantId = ctx.Session.TenentID

	err = r.cAdmin.ResourceNew(ctx.Session, res)
	r.WriteFinal(ctx.GinCtx, err)
}

func (r *R) ResourceGet(ctx request.Ctx) {
	resp, err := r.cAdmin.ResourceGet(ctx.Session, ctx.GinCtx.Param("slug"))
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) ResourceUpdate(ctx request.Ctx) {
	data := make(map[string]interface{})
	err := ctx.GinCtx.BindJSON(&data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}
	err = r.cAdmin.ResourceUpdate(ctx.Session, ctx.GinCtx.Param("slug"), data)
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) ResourceRemove(ctx request.Ctx) {
	err := r.cAdmin.ResourceDel(ctx.Session, ctx.GinCtx.Param("slug"))
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) ResourceList(ctx request.Ctx) {
	resp, err := r.cAdmin.ResourceList(ctx.Session)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) ResourceAgentList(ctx request.Ctx) {
	query := &vmodels.ResourceQuery{}
	err := ctx.GinCtx.BindJSON(&query)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	resp, err := r.cAdmin.ResourceAgentList(ctx.Session, query)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

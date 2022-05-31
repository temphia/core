package routes

import (
	"strconv"

	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/lib/apiutils/request"
)

func (r *R) UpdateTenant(ctx request.Ctx) {
	// ten := &entities.Tenant{}
	// if ten.Slug != cm.TenentID {
	// 	apiutils.WriteErr(c, ErrNotAllowd)
	// 	return
	// }
	// err := r.syncCtrl.UpdateTenant(ten)
	// apiutils.WriteFinal(c, err)
}

// domain

func (r *R) AddTenantDomain(ctx request.Ctx) {
	data := &entities.TenantDomain{}

	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.AddDomain(ctx.Session, data)
	r.WriteFinal(ctx.GinCtx, err)
}

func (r *R) ListTenantDomain(ctx request.Ctx) {
	resp, err := r.cAdmin.ListDomain(ctx.Session)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) GetTenantDomain(ctx request.Ctx) {
	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}
	resp, err := r.cAdmin.GetDomain(ctx.Session, id)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) UpdateTenantDomain(ctx request.Ctx) {
	data := map[string]interface{}{}
	err := ctx.GinCtx.BindJSON(&data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.UpdateDomain(ctx.Session, id, data)
	r.WriteFinal(ctx.GinCtx, err)
}

func (r *R) RemoveTenantDomain(ctx request.Ctx) {
	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.RemoveDomain(ctx.Session, id)
	r.WriteFinal(ctx.GinCtx, err)
}

// widget

func (r *R) AddDomainWidget(ctx request.Ctx) {
	did, err := strconv.ParseInt(ctx.GinCtx.Param("did"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}
	data := &entities.DomainWidget{}
	err = ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	data.DomainId = did
	err = r.cAdmin.AddDomainWidget(ctx.Session, data)
	r.WriteFinal(ctx.GinCtx, err)
}

func (r *R) UpdateDomainWidget(ctx request.Ctx) {
	wid, err := strconv.ParseInt(ctx.GinCtx.Param("wid"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}
	data := map[string]interface{}{}

	ctx.GinCtx.BindJSON(&data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.UpdateDomainWidget(ctx.Session, wid, data)
	r.WriteFinal(ctx.GinCtx, err)
}

func (r *R) GetDomainWidget(ctx request.Ctx) {
	wid, err := strconv.ParseInt(ctx.GinCtx.Param("wid"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	resp, err := r.cAdmin.GetDomainWidget(ctx.Session, wid)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) RemoveDomainWidget(ctx request.Ctx) {
	wid, err := strconv.ParseInt(ctx.GinCtx.Param("wid"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.RemoveDomainWidget(ctx.Session, wid)
	r.WriteFinal(ctx.GinCtx, err)
}

func (r *R) ListDomainWidget(ctx request.Ctx) {
	did, err := strconv.ParseInt(ctx.GinCtx.Param("did"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	resp, err := r.cAdmin.ListDomainWidget(ctx.Session, did)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

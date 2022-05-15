package routes

import (
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/lib/apiutils/request"
)

func (r *R) AddUserGroup(ctx request.Ctx) {
	group := &entities.UserGroup{}

	err := ctx.GinCtx.BindJSON(group)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	group.TenantID = ctx.Session.TenentID

	r.WriteJSON(ctx.GinCtx, nil, r.cAdmin.AddUserGroup(ctx.Session, group))
}

func (r *R) GetUserGroup(ctx request.Ctx) {
	resp, err := r.cAdmin.GetUserGroup(ctx.Session, ctx.GinCtx.Param("user_group"))
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) ListUserGroup(ctx request.Ctx) {
	resp, err := r.cAdmin.ListUserGroup(ctx.Session)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) UpdateUserGroup(ctx request.Ctx) {
	data := make(map[string]interface{})
	err := ctx.GinCtx.BindJSON(&data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.UpdateUserGroup(ctx.Session, ctx.GinCtx.Param("user_group"), data)
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) RemoveUserGroup(ctx request.Ctx) {
	r.WriteJSON(ctx.GinCtx, nil, r.cAdmin.RemoveUserGroup(ctx.Session, ctx.GinCtx.Param("user_group")))
}

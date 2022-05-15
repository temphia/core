package routes

import (
	"strconv"

	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/lib/apiutils/request"
)

// auth

func (r *R) AddUserGroupAuth(ctx request.Ctx) {
	data := &entities.UserGroupAuth{}
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.WriteJSON(
		ctx.GinCtx,
		nil,
		r.cAdmin.AddUserGroupAuth(ctx.Session, ctx.GinCtx.Param("ugroup"), data),
	)
}

func (r *R) UpdateUserGroupAuth(ctx request.Ctx) {
	data := make(map[string]interface{})
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.WriteJSON(
		ctx.GinCtx,
		nil,
		r.cAdmin.UpdateUserGroupAuth(ctx.Session, ctx.GinCtx.Param("ugroup"), id, data),
	)
}

func (r *R) ListUserGroupAuth(ctx request.Ctx) {
	data, err := r.cAdmin.ListUserGroupAuth(ctx.Session, ctx.GinCtx.Param("ugroup"))
	r.WriteJSON(ctx.GinCtx, data, err)
}

func (r *R) GetUserGroupAuth(ctx request.Ctx) {

	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupAuth(ctx.Session, ctx.GinCtx.Param("ugroup"), id)
	r.WriteJSON(ctx.GinCtx, resp, err)
}
func (r *R) RemoveUserGroupAuth(ctx request.Ctx) {
	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.RemoveUserGroupAuth(ctx.Session, ctx.GinCtx.Param("ugroup"), id)
	r.WriteJSON(ctx.GinCtx, nil, err)
}

// hook

func (r *R) AddUserGroupHook(ctx request.Ctx) {
	data := &entities.UserGroupHook{}
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.WriteJSON(
		ctx.GinCtx,
		nil,
		r.cAdmin.AddUserGroupHook(ctx.Session, ctx.GinCtx.Param("ugroup"), data),
	)
}

func (r *R) UpdateUserGroupHook(ctx request.Ctx) {
	data := make(map[string]interface{})
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.WriteJSON(
		ctx.GinCtx,
		nil,
		r.cAdmin.UpdateUserGroupHook(ctx.Session, ctx.GinCtx.Param("ugroup"), id, data),
	)
}

func (r *R) ListUserGroupHook(ctx request.Ctx) {
	resp, err := r.cAdmin.ListUserGroupHook(ctx.Session, ctx.GinCtx.Param("ugroup"))
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) GetUserGroupHook(ctx request.Ctx) {
	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupHook(ctx.Session, ctx.GinCtx.Param("ugroup"), id)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) RemoveUserGroupHook(ctx request.Ctx) {
	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.cAdmin.RemoveUserGroupHook(ctx.Session, ctx.GinCtx.Param("ugroup"), id)
}

// plug

func (r *R) AddUserGroupPlug(ctx request.Ctx) {
	data := &entities.UserGroupPlug{}
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.WriteJSON(
		ctx.GinCtx,
		nil,
		r.cAdmin.AddUserGroupPlug(ctx.Session, ctx.GinCtx.Param("ugroup"), data),
	)
}

func (r *R) UpdateUserGroupPlug(ctx request.Ctx) {
	data := make(map[string]interface{})
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.UpdateUserGroupPlug(ctx.Session, ctx.GinCtx.Param("ugroup"), id, data)
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) ListUserGroupPlug(ctx request.Ctx) {
	resp, err := r.cAdmin.ListUserGroupPlug(ctx.Session, ctx.GinCtx.Param("ugroup"))
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) GetUserGroupPlug(ctx request.Ctx) {
	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupPlug(ctx.Session, ctx.GinCtx.Param("ugroup"), id)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) RemoveUserGroupPlug(ctx request.Ctx) {
	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.cAdmin.RemoveUserGroupPlug(ctx.Session, ctx.GinCtx.Param("ugroup"), id)
}

// data

func (r *R) AddUserGroupData(ctx request.Ctx) {
	data := &entities.UserGroupData{}
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.WriteJSON(
		ctx.GinCtx,
		nil,
		r.cAdmin.AddUserGroupData(ctx.Session, ctx.GinCtx.Param("ugroup"), data),
	)
}

func (r *R) UpdateUserGroupData(ctx request.Ctx) {
	data := make(map[string]interface{})
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.UpdateUserGroupData(ctx.Session, ctx.GinCtx.Param("ugroup"), id, data)
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) ListUserGroupData(ctx request.Ctx) {
	resp, err := r.cAdmin.ListUserGroupData(ctx.Session, ctx.GinCtx.Param("ugroup"))
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) GetUserGroupData(ctx request.Ctx) {
	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupData(ctx.Session, ctx.GinCtx.Param("ugroup"), id)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) RemoveUserGroupData(ctx request.Ctx) {
	id, err := strconv.ParseInt(ctx.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.cAdmin.RemoveUserGroupData(ctx.Session, ctx.GinCtx.Param("ugroup"), id)
}

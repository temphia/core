package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/lib/apiutils"
	"github.com/temphia/core/backend/server/lib/apiutils/request"
)

func (r *R) AddUser(ctx request.Ctx) {
	usr := &entities.User{}

	err := ctx.GinCtx.BindJSON(usr)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}
	usr.TenantID = ctx.Session.TenentID

	r.WriteJSON(ctx.GinCtx, nil, r.cAdmin.AddUser(ctx.Session, usr))
}

func (r *R) UpdateUser(ctx request.Ctx) {
	data := make(map[string]interface{})

	err := ctx.GinCtx.BindJSON(&data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.WriteJSON(ctx.GinCtx, nil, r.cAdmin.UpdateUser(ctx.Session, data))
}

func (r *R) RemoveUser(ctx request.Ctx) {
	r.WriteJSON(ctx.GinCtx, nil, r.cAdmin.RemoveUser(ctx.Session, ctx.GinCtx.Param("user_id")))
}

func (r *R) GetUserByID(ctx request.Ctx) {
	usr, err := r.cAdmin.GetUserByID(ctx.Session, ctx.GinCtx.Param("user_id"))
	if usr != nil {
		usr.Password = ""
	}

	r.WriteJSON(ctx.GinCtx, usr, err)
}

func (r *R) ListUsers(ctx request.Ctx) {
	ugroup := ctx.GinCtx.Query("user_group")
	var usrs []*entities.User
	var err error

	if ugroup != "" {
		usrs, err = r.cAdmin.ListUsersByGroup(ctx.Session, ugroup)
	} else {
		usrs, err = r.cAdmin.ListUsers(ctx.Session)
	}

	r.WriteJSON(ctx.GinCtx, usrs, err)
}

func (r *R) UserProfileImage(ctx *gin.Context) {
	bytes, err := r.data.GetAsset("default_user_profile2.png")
	if err != nil {
		return
	}
	apiutils.WriteBinary(ctx, bytes)
}

// perm placeholder stuff

func (r *R) AddPerm(ctx request.Ctx) {

}
func (r *R) UpdatePerm(ctx request.Ctx) {

}
func (r *R) GetPerm(ctx request.Ctx) {

}
func (r *R) RemovePerm(ctx request.Ctx) {

}
func (r *R) AddRole(ctx request.Ctx) {

}
func (r *R) GetRole(ctx request.Ctx) {

}
func (r *R) UpdateRole(ctx request.Ctx) {

}
func (r *R) RemoveRole(ctx request.Ctx) {

}
func (r *R) AddUserRole(ctx request.Ctx) {

}
func (r *R) RemoveUserRole(ctx request.Ctx) {

}
func (r *R) ListAllPerm(ctx request.Ctx) {

}
func (r *R) ListAllRole(ctx request.Ctx) {

}
func (r *R) ListAllUserRole(ctx request.Ctx) {

}
func (r *R) ListAllUserPerm(ctx request.Ctx) {

}
func (r *R) ListUserPerm(ctx request.Ctx) {

}

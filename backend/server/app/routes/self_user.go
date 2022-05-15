package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/lib/apiutils"
	"github.com/temphia/core/backend/server/lib/apiutils/request"
	"github.com/temphia/core/backend/server/services/sockcore/transports"
)

type messageUser struct {
	UserId  string `json:"user_id,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *R) SelfMessageUser(req request.Ctx) {
	data := messageUser{}
	err := req.GinCtx.BindJSON(&data)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	_, err = r.cBasic.MessageUser(req.Session, data.UserId, data.Message)
	r.WriteFinal(req.GinCtx, err)
}

func (r *R) SelfGetUserInfo(req request.Ctx) {
	user := req.GinCtx.Param("user")
	resp, err := r.cBasic.GetUserInfo(req.Session, user)
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) SelfGetInfo(req request.Ctx) {
	resp, err := r.cBasic.GetSelfInfo(req.Session)
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) SelfModifyMessages(req request.Ctx) {
	opts := &entities.ModifyMessages{}
	err := req.GinCtx.BindJSON(opts)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cBasic.ModifyMessages(req.Session, opts)
	r.WriteFinal(req.GinCtx, err)
}

func (r *R) SelfListMessages(req request.Ctx) {
	opts := &entities.UserMessageReq{}
	err := req.GinCtx.BindJSON(opts)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cBasic.ListMessages(req.Session, opts)
	r.WriteJSON(req.GinCtx, resp, err)
}

type UserSocketUpdateOptions struct {
	Room       string   `json:"room,omitempty"`
	AddTags    []string `json:"add_tags,omitempty"`
	RemoveTags []string `json:"remove_tags,omitempty"`
	Clear      bool     `json:"clear,omitempty"`
}

func (r *R) UserSocketUpdate(ctx request.Ctx) {
	data := &UserSocketUpdateOptions{}
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.sockd.RoomUpdateTags(
		ctx.Session.TenentID,
		data.Room,
		ctx.Session.SessionID,
		&service.UpdateTagOptions{
			AddTags:    data.AddTags,
			ClearOld:   data.Clear,
			RemoveTags: data.RemoveTags,
		})

	r.WriteFinal(ctx.GinCtx, err)
}

type DgroupSockdChangeOptions struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	Ticket string `json:"ticket,omitempty"`
}

func (r *R) SockdDgroupChange(ctx request.Ctx) {
	data := &DgroupSockdChangeOptions{}
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.sockd.RoomUpdateTags(
		ctx.Session.TenentID,
		btypes.ROOM_SYSTABLE,
		ctx.Session.SessionID,
		&service.UpdateTagOptions{
			AddTags:    []string{fmt.Sprintf("dgroup.%s.%s", data.Source, data.Group)},
			ClearOld:   true,
			RemoveTags: []string{},
		})

	r.WriteFinal(ctx.GinCtx, err)
}

func (r *R) SelfUserSocket(c *gin.Context) {

	tenant := c.Param("tenant_id")
	token := c.Query("token")

	claim := claim.Session{}
	err := r.signer.Parse(tenant, token, &claim)
	if err != nil {
		apiutils.WriteErr(c, err.Error())
		return
	}

	// fixme => check claim proper type

	conn, err := transports.NewConnWS(c, claim.SessionID)
	if err != nil {
		apiutils.WriteErr(c, err.Error())
		return
	}

	connTags := []string{
		fmt.Sprint("sys.user_", claim.UserID),
		fmt.Sprint("sys.ugroup_", claim.UserGroup),
		fmt.Sprint("sys.device_", claim.DeviceID),
		btypes.TAG_REALUSER,
		btypes.TAG_CONSOLE_CONN,
	}

	err = r.sockd.NewConnection(&service.ConnOptions{
		NameSpace: tenant,
		Conn:      conn,
		Expiry:    10000,
		PreJoinRooms: map[string][]string{
			btypes.ROOM_SYS_USERS: connTags,
			btypes.ROOM_SYSTABLE:  connTags,
		},
	})
	if err != nil {
		apiutils.WriteErr(c, err.Error())
		return
	}

}

// fixme => impl placeholder

func (r *R) SelfChangeEmail(req request.Ctx) {

}

func (r *R) SelfChangeAuth(req request.Ctx) {

}

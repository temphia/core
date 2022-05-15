package routes

import (
	"io"

	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/models/vmodels"

	"github.com/temphia/temphia/backend/server/lib/apiutils"
	"github.com/temphia/temphia/backend/server/lib/apiutils/request"
)

func (r *R) BprintList(ctx request.Ctx) {

	rep, err := r.cAdmin.BprintList(ctx.Session, "")
	r.WriteJSON(ctx.GinCtx, rep, err)
}

func (r *R) BprintCreate(ctx request.Ctx) {
	bprint := &entities.BPrint{}
	err := ctx.GinCtx.BindJSON(bprint)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}
	id, err := r.cAdmin.BprintCreate(ctx.Session, bprint)
	r.WriteJSON(ctx.GinCtx, id, err)
}

func (r *R) BprintGet(ctx request.Ctx) {
	resp, err := r.cAdmin.BprintGet(ctx.Session, ctx.GinCtx.Param("id"))
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) BprintUpdate(ctx request.Ctx) {
	bprint := &entities.BPrint{}
	err := ctx.GinCtx.BindJSON(bprint)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}
	err = r.cAdmin.BprintUpdate(ctx.Session, bprint)
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) BprintRemove(ctx request.Ctx) {
	err := r.cAdmin.BprintRemove(ctx.Session, ctx.GinCtx.Param("id"))
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) BprintInstall(ctx request.Ctx) {
	opts := &vmodels.RepoInstallOpts{}
	err := ctx.GinCtx.BindJSON(opts)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	resp, err := r.cAdmin.BprintInstall(ctx.Session, opts)
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (c *R) BprintImport(ctx request.Ctx) {

	opts := &vmodels.RepoImportOpts{}
	err := ctx.GinCtx.BindJSON(opts)
	if err != nil {
		return
	}

	resp, err := c.cAdmin.BprintImport(ctx.Session, opts)

	c.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) BprintListFiles(ctx request.Ctx) {
	resp, err := r.cAdmin.BprintListBlobs(ctx.Session, ctx.GinCtx.Param("id"))
	r.WriteJSON(ctx.GinCtx, resp, err)
}

func (r *R) BprintNewBlob(ctx request.Ctx) {
	bytes, err := readForm(ctx.GinCtx)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.BprintNewBlob(ctx.Session, ctx.GinCtx.Param("id"), ctx.GinCtx.Param("file_id"), bytes)
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) BprintUpdateBlob(ctx request.Ctx) {
	bytes, err := io.ReadAll(ctx.GinCtx.Request.Body)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	err = r.cAdmin.BprintUpdateBlob(ctx.Session, ctx.GinCtx.Param("id"), ctx.GinCtx.Param("file_id"), bytes)
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) BprintGetFile(ctx request.Ctx) {
	out, err := r.cAdmin.BprintGetBlob(ctx.Session, ctx.GinCtx.Param("id"), ctx.GinCtx.Param("file_id"))
	if err != nil {
		return
	}

	apiutils.WriteBinary(ctx.GinCtx, out)
}

func (r *R) BprintDelFile(ctx request.Ctx) {
	err := r.cAdmin.BprintDeleteBlob(ctx.Session, ctx.GinCtx.Param("id"), ctx.GinCtx.Param("file_id"))
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) BprintPushToken(ctx request.Ctx) {

}

package routes

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/lib/apiutils"
	"github.com/temphia/temphia/backend/server/lib/apiutils/request"
)

func (r *R) ListCabinetSources(req request.Ctx) {

	sources, err := r.cBasic.ListCabinetSources(req.Session)
	r.WriteJSON(req.GinCtx, sources, err)
}

func (r *R) NewFolder(req request.Ctx) {
	r.WriteFinal(
		req.GinCtx,
		r.cCabinet.AddFolder(req.Session, req.GinCtx.Param("folder")),
	)
}

func (r *R) UploadFile(req request.Ctx) {
	bytes, err := readForm(req.GinCtx)

	err = r.cCabinet.AddBlob(req.Session, req.GinCtx.Param("folder"), req.GinCtx.Param("fname"), bytes)
	r.WriteFinal(req.GinCtx, err)
}

func (r *R) ListRootFolder(req request.Ctx) {
	folders, err := r.cCabinet.ListRoot(req.Session)
	r.WriteJSON(req.GinCtx, folders, err)
}

func (r *R) ListFolder(req request.Ctx) {
	files, err := r.cCabinet.ListFolder(req.Session, req.GinCtx.Param("folder"))
	r.WriteJSON(req.GinCtx, files, err)
}

func (r *R) GetFile(req request.Ctx) {

	bytes, err := r.cCabinet.GetBlob(req.Session, req.GinCtx.Param("folder"), req.GinCtx.Param("fname"))
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	req.GinCtx.Writer.WriteHeader(http.StatusOK)
	req.GinCtx.Writer.Write(bytes)
}

func (r *R) DeleteFile(req request.Ctx) {
	err := r.cCabinet.DeleteBlob(req.Session, req.GinCtx.Param("folder"), req.GinCtx.Param("fname"))
	r.WriteFinal(req.GinCtx, err)
}

func (r *R) GetFolderTicket(req request.Ctx) {
	resp, err := r.cCabinet.NewFolderTicket(req.Session, req.GinCtx.Param("folder"))
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) GetFilePreview(req request.Ctx) {
	// bytes, err := b.blobfs.GetBlobPreview(c.Request.Context(), cm.TenentID, c.Param("folder"), c.Param("fname"))
	// if err != nil {
	// 	apiutils.WriteErr(c, err.Error())
	// 	return
	// }
	// c.Writer.Header().Add("Content-Type", "application/octet-stream")
	// c.Writer.Write(bytes)
}

func (r *R) TicketCabinetFile(ctx *gin.Context) {
	ticket := ctx.Param("ticket")
	tenantId := ctx.Param("tenant_id")
	file := ctx.Param("file")

	ct := &claim.TicketCabinet{}
	err := r.signer.Parse(tenantId, ticket, ct)
	if err != nil {
		r.WriteErr(ctx, err.Error())
		return
	}

	// fixme => check ct.DeviceId

	out, err := r.cCabinet.TicketFile(tenantId, file, ct)
	if err != nil {
		r.WriteErr(ctx, err.Error())
		return
	}

	apiutils.WriteBinary(ctx, out)
}

func (r *R) TicketCabinetPreviewFile(ctx *gin.Context) {
	ticket := ctx.Param("ticket")
	tenantId := ctx.Param("tenant_id")
	file := ctx.Param("file")

	ct := &claim.TicketCabinet{}
	err := r.signer.Parse(tenantId, ticket, ct)
	if err != nil {
		r.WriteErr(ctx, err.Error())
		return
	}
	out, err := r.cCabinet.TicketFile(tenantId, file, ct)
	if err != nil {
		r.WriteErr(ctx, err.Error())
		return
	}

	apiutils.WriteBinary(ctx, out)
}

func (r *R) TicketCabinetList(ctx *gin.Context) {
	ticket := ctx.Param("ticket")
	tenantId := ctx.Param("tenant_id")

	ct := &claim.TicketCabinet{}
	err := r.signer.Parse(tenantId, ticket, ct)
	if err != nil {
		r.WriteErr(ctx, err.Error())
		return
	}
	resp, err := r.cCabinet.TicketList(tenantId, ct)
	if err != nil {
		r.WriteErr(ctx, err.Error())
		return
	}

	r.WriteJSON(ctx, resp, err)
}

func (r *R) TicketCabinetUpload(ctx *gin.Context) {
	ticket := ctx.Param("ticket")
	tenantId := ctx.Param("tenant_id")
	file := ctx.Param("file")

	ct := &claim.TicketCabinet{}
	err := r.signer.Parse(tenantId, ticket, ct)
	if err != nil {
		r.WriteErr(ctx, err.Error())
		return
	}

	out, err := readForm(ctx)
	if err != nil {
		r.WriteErr(ctx, err.Error())
		return
	}

	err = r.cCabinet.TicketUpload(tenantId, file, out, ct)
	r.WriteFinal(ctx, err)
}

func readForm(ctx *gin.Context) ([]byte, error) {
	fh, err := ctx.FormFile("file")
	if err != nil {
		return nil, err
	}

	file, err := fh.Open()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}

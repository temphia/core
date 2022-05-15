package routes

import (
	"strconv"

	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/temphia/core/backend/server/lib/apiutils/request"
)

type newRowReq struct {
	Cells map[string]interface{} `json:"cells,omitempty"`
}

func (r *R) NewRow(req request.Ctx) {

	data := &newRowReq{}
	err := req.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	id, err := r.cDtable.NewRow(req.Session, req.GinCtx.Param("table_id"), data.Cells)
	r.WriteJSON(req.GinCtx, id, err)
}

func (r *R) GetRow(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	cells, err := r.cDtable.GetRow(req.Session, req.GinCtx.Param("table_id"), id)
	r.WriteJSON(req.GinCtx, cells, err)
}

type updateRowReq struct {
	Version int64                  `json:"version,omitempty"`
	Cells   map[string]interface{} `json:"cells,omitempty"`
}

func (r *R) UpdateRow(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	data := &updateRowReq{}
	err = req.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	cells, err := r.cDtable.UpdateRow(req.Session, req.GinCtx.Param("table_id"), id, data.Version, data.Cells)
	r.WriteJSON(req.GinCtx, cells, err)
}

func (r *R) DeleteRow(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.DeleteRow(req.Session, req.GinCtx.Param("table_id"), id)
	r.WriteFinal(req.GinCtx, err)
}

func (r *R) SimpleQuery(req request.Ctx) {
	query := store.SimpleQueryReq{}
	err := req.GinCtx.BindJSON(&query)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.SimpleQuery(req.Session, req.GinCtx.Param("table_id"), query)
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) FTSQuery(req request.Ctx) {
	query := store.FTSQueryReq{}
	err := req.GinCtx.BindJSON(&query)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.FTSQuery(req.Session, req.GinCtx.Param("table_id"), query.SearchTerm)
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) RefLoad(req request.Ctx) {
	query := &store.RefLoadReq{}
	err := req.GinCtx.BindJSON(&query)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.RefLoad(req.Session, query)
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) RefResolve(req request.Ctx) {
	query := &store.RefResolveReq{}
	err := req.GinCtx.BindJSON(&query)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.RefResolve(req.Session, query)
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) ReverseRefLoad(req request.Ctx) {
	query := &store.RevRefLoadReq{}

	err := req.GinCtx.BindJSON(query)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.ReverseRefLoad(req.Session, query)
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) ListActivity(req request.Ctx) {
	rid, err := strconv.ParseInt(req.GinCtx.Param("row_id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.ListActivity(req.Session, req.GinCtx.Param("table_id"), int(rid))
	r.WriteJSON(req.GinCtx, resp, err)
}

type commentRowReq struct {
	Message string `json:"message,omitempty"`
}

func (r *R) CommentRow(req request.Ctx) {
	rid, err := strconv.ParseInt(req.GinCtx.Param("row_id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	reqdata := commentRowReq{}

	err = req.GinCtx.BindJSON(&reqdata)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.CommentRow(req.Session, req.GinCtx.Param("table_id"), reqdata.Message, int(rid))
	r.WriteFinal(req.GinCtx, err)
}

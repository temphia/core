package routes

import (
	"strconv"

	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/lib/apiutils/request"
)

func (r *R) ListDtableSources(req request.Ctx) {
	sources, err := r.cBasic.ListDyndbSources(req.Session)
	r.WriteJSON(req.GinCtx, sources, err)
}

// dyn_table_group

func (r *R) NewGroup(req request.Ctx) {
	tg := &entities.NewTableGroup{}
	err := req.GinCtx.BindJSON(tg)

	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.NewGroup(req.Session, req.GinCtx.Param("source"), tg)
	r.WriteFinal(req.GinCtx, err)

}
func (r *R) EditGroup(req request.Ctx) {
	tg := &entities.TableGroupPartial{}
	err := req.GinCtx.BindJSON(tg)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.EditGroup(req.Session, req.GinCtx.Param("source"), req.GinCtx.Param("group_id"), tg)
	r.WriteFinal(req.GinCtx, err)
}

func (r *R) GetGroup(req request.Ctx) {
	resp, err := r.cDtable.GetGroup(req.Session, req.GinCtx.Param("source"), req.GinCtx.Param("group_id"))
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) ListGroup(req request.Ctx) {
	gr, err := r.cDtable.ListGroup(req.Session, req.GinCtx.Param("source"))
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	r.WriteJSON(req.GinCtx, gr, err)
}

func (r *R) LoadGroup(req request.Ctx) {
	gr, err := r.cDtable.LoadGroup(req.Session)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	r.WriteJSON(req.GinCtx, gr, err)
}

func (r *R) DeleteGroup(req request.Ctx) {
	err := r.cDtable.DeleteGroup(req.Session, req.GinCtx.Param("source"), req.GinCtx.Param("group_id"))
	r.WriteFinal(req.GinCtx, err)

}

// dyn_table

func (r *R) AddTable(req request.Ctx) {
	t := &entities.NewTable{}
	err := req.GinCtx.BindJSON(t)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.AddTable(req.Session, t)
	r.WriteFinal(req.GinCtx, err)

}
func (r *R) EditTable(req request.Ctx) {
	tp := &entities.TablePartial{}
	err := req.GinCtx.BindJSON(tp)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.EditTable(req.Session, req.GinCtx.Param("table_id"), tp)
	r.WriteFinal(req.GinCtx, err)

}

func (r *R) GetTable(req request.Ctx) {
	tbl, err := r.cDtable.GetTable(req.Session, req.GinCtx.Param("table_id"))
	r.WriteJSON(req.GinCtx, tbl, err)
}

func (r *R) ListTables(req request.Ctx) {
	tbls, err := r.cDtable.ListTables(req.Session)
	r.WriteJSON(req.GinCtx, tbls, err)

}
func (r *R) DeleteTable(req request.Ctx) {
	err := r.cDtable.DeleteTable(req.Session, req.GinCtx.Param("table_id"))
	r.WriteFinal(req.GinCtx, err)

}

// dyn_table_column

func (r *R) AddColumn(req request.Ctx) {
	nc := &entities.NewColumn{}
	err := req.GinCtx.BindJSON(nc)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.AddColumn(req.Session, req.GinCtx.Param("table_id"), nc)
	r.WriteFinal(req.GinCtx, err)

}
func (r *R) EditColumn(req request.Ctx) {
	cp := &entities.ColumnPartial{}
	err := req.GinCtx.BindJSON(cp)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.EditColumn(req.Session, req.GinCtx.Param("table_id"), req.GinCtx.Param("column_id"), cp)
	r.WriteFinal(req.GinCtx, err)

}

func (r *R) GetColumn(req request.Ctx) {
	resp, err := r.cDtable.GetColumn(req.Session, req.GinCtx.Param("table_id"), req.GinCtx.Param("column_id"))
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) ListColumns(req request.Ctx) {
	cols, err := r.cDtable.ListColumns(req.Session, req.GinCtx.Param("table_id"))
	r.WriteJSON(req.GinCtx, cols, err)

}
func (r *R) DeleteColumn(req request.Ctx) {

	err := r.cDtable.DeleteColumn(req.Session, req.GinCtx.Param("table_id"), req.GinCtx.Param("column_id"))
	r.WriteFinal(req.GinCtx, err)

}
func (r *R) AddIndex(req request.Ctx) {

}

// dyn_table_meta

func (r *R) AddUniqueIndex(req request.Ctx) {

}
func (r *R) AddFTSIndex(req request.Ctx) {

}
func (r *R) AddColumnFRef(req request.Ctx) {

}
func (r *R) ListIndex(req request.Ctx) {

}
func (r *R) RemoveIndex(req request.Ctx) {

}

// view stuff

func (r *R) NewView(req request.Ctx) {
	view := entities.DataView{}
	err := req.GinCtx.BindJSON(&view)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.NewView(req.Session, req.GinCtx.Param("table_id"), &view)
	r.WriteFinal(req.GinCtx, err)
}

func (r *R) ModifyView(req request.Ctx) {
	view := make(map[string]interface{})

	err := req.GinCtx.BindJSON(&view)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.ModifyView(req.Session, req.GinCtx.Param("table_id"), id, view)
	r.WriteFinal(req.GinCtx, err)

}

func (r *R) GetView(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.GetView(req.Session, req.GinCtx.Param("table_id"), id)
	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) ListView(req request.Ctx) {
	resp, err := r.cDtable.ListView(req.Session, req.GinCtx.Param("table_id"))
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) DelView(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.DelView(req.Session, req.GinCtx.Param("table_id"), id)
	r.WriteFinal(req.GinCtx, err)
}

// hooks

func (r *R) NewHook(req request.Ctx) {
	hook := entities.DataHook{}
	err := req.GinCtx.BindJSON(&hook)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.NewHook(req.Session, req.GinCtx.Param("table_id"), &hook)
	r.WriteFinal(req.GinCtx, err)
}

func (r *R) ModifyHook(req request.Ctx) {
	data := make(map[string]interface{})

	err := req.GinCtx.BindJSON(&data)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.ModifyHook(req.Session, req.GinCtx.Param("table_id"), id, data)
	r.WriteFinal(req.GinCtx, err)

}

func (r *R) GetHook(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.GetHook(req.Session, req.GinCtx.Param("table_id"), id)
	r.WriteJSON(req.GinCtx, resp, err)

}

func (r *R) ListHook(req request.Ctx) {
	resp, err := r.cDtable.ListHook(req.Session, req.GinCtx.Param("table_id"))
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	r.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) DelHook(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.DelHook(req.Session, req.GinCtx.Param("table_id"), id)
	r.WriteFinal(req.GinCtx, err)
}

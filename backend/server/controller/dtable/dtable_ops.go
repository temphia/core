package dtable

import (
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/store"
)

func getTarget(uclaim *claim.Session) (string, string) {
	return uclaim.ServicePath[1], uclaim.ServicePath[2]
}

func (d *Controller) NewRow(uclaim *claim.Session, tslug string, cells map[string]interface{}) (int64, error) {

	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)

	return dynDb.NewRow(0, store.NewRowReq{
		TenantId: uclaim.TenentID,
		Group:    group,
		Table:    tslug,
		Data:     cells,
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

func (d *Controller) GetRow(uclaim *claim.Session, tslug string, id int64) (map[string]interface{}, error) {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)

	return dynDb.GetRow(0, store.GetRowReq{
		TenantId: uclaim.TenentID,
		Group:    group,
		Table:    tslug,
		Id:       id,
	})
}

func (d *Controller) UpdateRow(uclaim *claim.Session, tslug string, id, version int64, cells map[string]interface{}) (map[string]interface{}, error) {

	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)
	return dynDb.UpdateRow(0, store.UpdateRowReq{
		TenantId: uclaim.TenentID,
		Version:  version,
		Group:    group,
		Table:    tslug,
		Data:     cells,
		Id:       id,
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

func (d *Controller) DeleteRow(uclaim *claim.Session, tslug string, id int64) error {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)
	return dynDb.DeleteRows(0, store.DeleteRowReq{
		TenantId: uclaim.TenentID,
		Group:    group,
		Table:    tslug,
		Id:       []int64{id},
	})
}

func (d *Controller) SimpleQuery(uclaim *claim.Session, tslug string, query store.SimpleQueryReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)

	query.TenantId = uclaim.TenentID
	query.Table = tslug
	query.Group = group

	return dynDb.SimpleQuery(0, query)
}

func (d *Controller) FTSQuery(uclaim *claim.Session, tslug, qstr string) (*store.QueryResult, error) {

	source, group := getTarget(uclaim)

	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)
	return dynDb.FTSQuery(0, store.FTSQueryReq{
		TenantId:   uclaim.TenentID,
		Table:      tslug,
		Group:      group,
		SearchTerm: qstr,
		Count:      10,
	})
}

func (d *Controller) TemplateQuery(uclaim *claim.Session, tslug string, query interface{}) (*store.QueryResult, error) {

	// source, group := getTarget(uclaim)

	// dynDb := d.hub.GetSource(source, uclaim.TenentID)
	// // fixme
	// return dynDb.TemplateQuery(0, store.TemplateQueryReq{
	// 	TenantId:  uclaim.TenentID,
	// 	Group:     group,
	// 	Fragments: nil,
	// 	Name:      "",
	// })

	return nil, nil
}

func (d *Controller) RefResolve(uclaim *claim.Session, req *store.RefResolveReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)
	return dynDb.RefResolve(0, group, req)
}

func (d *Controller) ReverseRefLoad(uclaim *claim.Session, req *store.RevRefLoadReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)
	return dynDb.ReverseRefLoad(0, group, req)
}

func (d *Controller) RefLoad(uclaim *claim.Session, req *store.RefLoadReq) (*store.QueryResult, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)
	return dynDb.RefLoad(0, group, req)
}

func (d *Controller) ListActivity(uclaim *claim.Session, table string, rowId int) ([]*entities.DynActivity, error) {
	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)
	return dynDb.ListActivity(group, table, rowId)
}

func (d *Controller) CommentRow(uclaim *claim.Session, table, msg string, rowId int) error {

	source, group := getTarget(uclaim)
	dynDb := d.dynHub.GetSource(source, uclaim.TenentID)

	return dynDb.NewActivity(group, table, &entities.DynActivity{
		Type:      "comment",
		RowId:     int64(rowId),
		RowVerson: 0,
		UserId:    uclaim.UserID,
		Payload:   msg,
	})
}

package store

import (
	"encoding/json"

	"github.com/temphia/core/backend/server/btypes/models/entities"
)

type ModCtx struct {
	UserId    string `json:"user_id,omitempty"`
	UserSign  string `json:"user_sign,omitempty"`
	InitSign  string `json:"init_sign,omitempty"`
	TableName string `json:"table_name,omitempty"`
}

func (m *ModCtx) JSON() ([]byte, error) {
	return json.Marshal(m)
}

type NewRowReq struct {
	TenantId string                 `json:"-"`
	Group    string                 `json:"group,omitempty"`
	Table    string                 `json:"table,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
	ModCtx   ModCtx                 `json:"mod_ctx,omitempty"`
}

type NewBatchRowReq struct {
	TenantId string                   `json:"-"`
	Group    string                   `json:"group,omitempty"`
	Table    string                   `json:"table,omitempty"`
	Data     []map[string]interface{} `json:"data,omitempty"`
	ModCtx   ModCtx                   `json:"mod_ctx,omitempty"`
}

type GetRowReq struct {
	TenantId  string `json:"-"`
	Group     string `json:"group,omitempty"`
	Table     string `json:"table,omitempty"`
	Id        int64  `json:"id,omitempty"`
	SkipCache bool   `json:"skip_cache,omitempty"`
}

type UpdateRowReq struct {
	TenantId string                 `json:"-"`
	Id       int64                  `json:"id,omitempty"`
	Version  int64                  `json:"version,omitempty"`
	Group    string                 `json:"group,omitempty"`
	Table    string                 `json:"table,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
	ModCtx   ModCtx                 `json:"mod_ctx,omitempty"`
}

type DeleteRowReq struct {
	TenantId string  `json:"-"`
	Group    string  `json:"group,omitempty"`
	Table    string  `json:"table,omitempty"`
	Id       []int64 `json:"id,omitempty"`
	ModCtx   ModCtx  `json:"mod_ctx,omitempty"`
}

type SimpleQueryReq struct {
	TenantId      string        `json:"-"`
	Table         string        `json:"table,omitempty"`
	Group         string        `json:"group,omitempty"`
	Count         int64         `json:"count,omitempty"`
	FilterConds   []*FilterCond `json:"filter_conds,omitempty"`
	Page          int64         `json:"page,omitempty"`
	Selects       []string      `json:"selects,omitempty"`
	SearchTerm    string        `json:"search_term,omitempty"`
	LoadExtraMeta bool          `json:"load_extra_meta,omitempty"`
}

type FTSQueryReq struct {
	TenantId   string `json:"-"`
	Table      string `json:"table,omitempty"`
	Group      string `json:"group,omitempty"`
	SearchTerm string `json:"search_term,omitempty"`
	Count      int64  `json:"count,omitempty"`
}

type RefLoadReq struct {
	Column      string `json:"column,omitempty"`
	Type        string `json:"type,omitempty"`
	Target      string `json:"target,omitempty"`
	Object      string `json:"object,omitempty"`
	CursorRowId int64  `json:"cursor_row_id,omitempty"`
}

type RevRefLoadReq struct {
	CurrentTable string      `json:"current_table,omitempty"`
	TargetTable  string      `json:"target_table,omitempty"`
	TargetColumn string      `json:"column,omitempty"`
	CurrentItem  interface{} `json:"current_item,omitempty"`
	CursorRowId  int64       `json:"cursor_row_id,omitempty"`
	Count        int         `json:"count,omitempty"`
}

type RefResolveReq struct {
	Column string        `json:"column,omitempty"`
	Type   string        `json:"type,omitempty"`
	Target string        `json:"target,omitempty"`
	Object string        `json:"object,omitempty"`
	RowIds []interface{} `json:"row_ids,omitempty"`
}

type FilterCond struct {
	Column string      `json:"column,omitempty"`
	Cond   string      `json:"cond,omitempty"`
	Value  interface{} `json:"value,omitempty"`
	Target string      `json:"target,omitempty"`
}

type TemplateQueryReq struct {
	TenantId  string                            `json:"-"`
	Group     string                            `json:"group,omitempty"`
	Fragments map[string]map[string]interface{} `json:"fragments,omitempty"`
	Name      string                            `json:"name,omitempty"`
}

type RawQueryReq struct {
	TenantId    string `json:"-"`
	NoTransform string `json:"no_transform,omitempty"`
	Group       string `json:"group,omitempty"`
	QStr        string `json:"qstr,omitempty"`
}

type QueryResult struct {
	Count     int64                       `json:"count,omitempty"`
	Page      int64                       `json:"page,omitempty"`
	Rows      []map[string]interface{}    `json:"rows"`
	Columns   map[string]*entities.Column `json:"columns,omitempty"`
	ExtraMeta *QueryMeta                  `json:"extra_meta,omitempty"`
}

type QueryMeta struct {
	ReverseRefs []*entities.Column   `json:"reverse_refs,omitempty"`
	Hooks       []*entities.DataHook `json:"hooks,omitempty"`
	Views       []*entities.DataView `json:"views,omitempty"`
}

type LoadDgroupResp struct {
	Tables          []*entities.Table `json:"tables,omitempty"`
	CabinetTicket   string            `json:"cabinet_ticket,omitempty"`
	SockdRoomTicket string            `json:"sockd_room_ticket,omitempty"`
}

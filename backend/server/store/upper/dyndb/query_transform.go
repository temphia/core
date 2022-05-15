package dyndb

import (
	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/upper/db/v4"
)

const (
	filterEqual    = "equal"
	filterNotEqual = "not_equal"
	filterIn       = "in"
	filterNotIn    = "not_in"
	filterLT       = "less_than"
	filterGT       = "greater_than"
	filterLTE      = "less_than_or_equal"
	filterGTE      = "greater_than_or_equal"

	OptrEqual    = ""
	OptrNotEqual = " !="
	OptrIn       = " IN"
	OptrNotIn    = " NOT IN"
	OptrLT       = " <"
	OptrGT       = " >"
	OptrLTE      = " <="
	OptrGTE      = " >="
)

/*
	filterIsNULL = "is_null"
	filterLIKE   = "like"
	around(50m)
	not_around(50m)
	ref(target) => join
*/

var (
	OptrMap = map[string]string{
		filterEqual:    OptrEqual,
		filterNotEqual: OptrNotEqual,
		filterIn:       OptrIn,
		filterNotIn:    OptrNotIn,
		filterLT:       OptrLT,
		filterGT:       OptrGT,
		filterLTE:      OptrLTE,
		filterGTE:      OptrGTE,
	}
)

func (d *DynDB) transformFilters(fcs []*store.FilterCond) (db.Cond, error) {
	conds := make(db.Cond)

	for _, filter := range fcs {

		optr, ok := OptrMap[filter.Cond]
		if !ok {
			return conds, nil
		}
		conds[filter.Column+optr] = filter.Value
	}

	// conds["deleted_at"] = nil

	return conds, nil
}

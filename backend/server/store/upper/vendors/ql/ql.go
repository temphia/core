package ql

import (
	"database/sql"

	"github.com/temphia/core/backend/server/app/config"
	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/temphia/core/backend/server/lib/dbutils"
	"github.com/temphia/core/backend/server/registry"
	"github.com/temphia/core/backend/server/store/upper"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/ql"
)

func init() {
	registry.SetStoreBuilders(store.VendorQL, func(so *config.StoreSource) (store.Store, error) {
		return upper.NewAdapter(&qldb{})(so)
	})
}

type qldb struct{}

func (qldb) Db(conf *config.StoreSource) (db.Session, error) {
	var settings = ql.ConnectionURL{
		Database: conf.Name,
	}

	return ql.Open(settings)
}

func (qldb) NewTx(sqlTx *sql.Tx) (dbutils.Tx, error) {
	return ql.NewTx(sqlTx)
}

func (qldb) CoreSchema() string {
	return ""
}
func (qldb) DtableSchema() string {
	return ""
}

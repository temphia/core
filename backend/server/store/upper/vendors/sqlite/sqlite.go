package sqlite

import (
	"database/sql"

	"github.com/temphia/temphia/backend/server/app/config"
	"github.com/temphia/temphia/backend/server/btypes/store"
	"github.com/temphia/temphia/backend/server/lib/dbutils"
	"github.com/temphia/temphia/backend/server/registry"
	"github.com/temphia/temphia/backend/server/store/upper"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

func init() {
	registry.SetStoreBuilders(store.VendorSqlite, func(so *config.StoreSource) (store.Store, error) {
		return upper.NewAdapter(&sl{})(so)
	})
}

type sl struct{}

func (sl) Db(so *config.StoreSource) (db.Session, error) {
	var settings = sqlite.ConnectionURL{
		Database: so.Name,
	}

	return sqlite.Open(settings)
}

func (sl) NewTx(sqlTx *sql.Tx) (dbutils.Tx, error) {
	return sqlite.NewTx(sqlTx)
}

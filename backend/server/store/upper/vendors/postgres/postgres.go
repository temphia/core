package postgres

import (
	"database/sql"

	"github.com/temphia/temphia/backend/server/app/config"
	"github.com/temphia/temphia/backend/server/btypes/store"
	"github.com/temphia/temphia/backend/server/lib/dbutils"
	"github.com/temphia/temphia/backend/server/registry"
	"github.com/temphia/temphia/backend/server/store/upper"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

func init() {
	registry.SetStoreBuilders(store.VendorPostgres, func(so *config.StoreSource) (store.Store, error) {
		return upper.NewAdapter(&pq{})(so)
	})
}

type pq struct{}

func (pq) Db(conf *config.StoreSource) (db.Session, error) {
	var settings = postgresql.ConnectionURL{
		Database: conf.Target,
		User:     conf.User,
		Password: conf.Password,
		Socket:   "",
		Options:  make(map[string]string),
		Host:     conf.HostPath + ":" + conf.Port,
	}

	return postgresql.Open(settings)
}

func (pq) NewTx(sqlTx *sql.Tx) (dbutils.Tx, error) {
	return postgresql.NewTx(sqlTx)
}

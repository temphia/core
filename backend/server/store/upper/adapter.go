package upper

import (
	"github.com/temphia/core/backend/server/app/config"
	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/temphia/core/backend/server/lib/dbutils"
	"github.com/temphia/core/backend/server/lib/dlock"
	"github.com/temphia/core/backend/server/store/upper/coredb"
	"github.com/temphia/core/backend/server/store/upper/dyndb"
	"github.com/temphia/core/backend/server/store/upper/plugkv"
	"github.com/temphia/core/backend/server/store/upper/tns"
	"github.com/temphia/core/backend/server/store/upper/ucore"
	"github.com/temphia/core/backend/server/store/upper/zenerator"

	"github.com/upper/db/v4"
)

// Adapter is adapter for upper bashed implementation
type Adapter struct {
	db           db.Session
	uvendor      ucore.UpperVendor
	innerCoreDB  store.CoreDB
	innerStateDb store.PlugStateKV
	innerDynDB   store.DynDB
}

func NewAdapter(upvendor ucore.UpperVendor) func(conf *config.StoreSource) (store.Store, error) {

	return func(conf *config.StoreSource) (store.Store, error) {

		_tns := tns.New("shared")
		ztr := zenerator.New(conf.Vendor, _tns)

		sess, err := upvendor.Db(conf)
		if err != nil {
			return nil, err
		}

		return &Adapter{
			db:           sess,
			uvendor:      upvendor,
			innerCoreDB:  coredb.New(sess, conf.Vendor),
			innerStateDb: plugkv.New(sess, dbutils.NewTxMgr(upvendor.NewTx), conf.Vendor),
			innerDynDB: dyndb.New(ucore.DynDBOptions{
				Session:    sess,
				TxnManager: dbutils.NewTxMgr(upvendor.NewTx),
				DynGen:     ztr,
				TNS:        _tns,
				SharedLock: dlock.New(""),
			}),
		}, nil
	}

}

func (u *Adapter) Name() string {
	return ""
}

func (u *Adapter) Supports(dbType store.StoreType) bool {
	switch dbType {
	case store.TypeCoreDB, store.TypeStateDB:
		return true
	case store.TypeDynDB:
		return true
	case store.TypeBlobStore:
		return false
	default:
		return false
	}
}

func (u *Adapter) Migrate() error {
	return nil
}

func (u *Adapter) CoreDB() store.CoreDB {
	return u.innerCoreDB
}

func (u *Adapter) SyncDB() store.SyncDB {
	return nil
}

func (u *Adapter) StateDB() store.PlugStateKV {
	return u.innerStateDb
}

func (u *Adapter) DynDB() store.DynDB {
	return u.innerDynDB
}

func (u *Adapter) CabinetSource() store.CabinetSource {
	return nil
}

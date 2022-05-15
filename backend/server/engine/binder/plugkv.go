package binder

import (
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/thoas/go-funk"
)

type plugKvBinder struct {
	core    *Binder
	stateKv store.PlugStateKV
	txns    []uint32
}

func (pkv *plugKvBinder) checkTxn(tx uint32) error {
	if tx == 0 {
		return nil
	}
	if !funk.ContainsUInt32(pkv.txns, tx) {
		return easyerr.NotFound()
	}
	return nil
}

func (pkv *plugKvBinder) NewTxn() (uint32, error) {
	tx, err := pkv.stateKv.NewTxn()
	if err != nil {
		return 0, err
	}
	pkv.txns = append(pkv.txns, tx)
	return tx, nil
}

func (pkv *plugKvBinder) RollBack(txid uint32) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}

	return pkv.stateKv.RollBack(txid)
}

func (pkv *plugKvBinder) Commit(txid uint32) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}

	return pkv.stateKv.Commit(txid)
}

func (pkv *plugKvBinder) Set(txid uint32, key, value string, opts *store.SetOptions) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.Set(txid, pkv.core.namespace, pkv.core.plugId, key, value, opts)
}

func (pkv *plugKvBinder) Update(txid uint32, key, value string, opts *store.UpdateOptions) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.Update(txid, pkv.core.namespace, pkv.core.plugId, key, value, opts)
}

func (pkv *plugKvBinder) Get(txid uint32, key string) (*entities.PlugKV, error) {
	err := pkv.checkTxn(txid)
	if err != nil {
		return nil, err
	}
	return pkv.stateKv.Get(txid, pkv.core.namespace, pkv.core.plugId, key)
}

func (pkv *plugKvBinder) Del(txid uint32, key string) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.Del(txid, pkv.core.namespace, pkv.core.plugId, key)
}

func (pkv *plugKvBinder) DelBatch(txid uint32, keys []string) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.DelBatch(txid, pkv.core.namespace, pkv.core.plugId, keys)
}

func (pkv *plugKvBinder) Query(txid uint32, query *store.PkvQuery) ([]*entities.PlugKV, error) {
	err := pkv.checkTxn(txid)
	if err != nil {
		return nil, err
	}

	return pkv.stateKv.Query(txid, pkv.core.namespace, pkv.core.plugId, query)
}

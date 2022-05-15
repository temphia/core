package mbindings

import (
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/store"
)

func (m *MockedBindings) Set(txid uint32, key, value string, opts *store.SetOptions) error {
	if key == "err_one" {
		return easyerr.Error("set err")
	}
	return nil
}

func (m *MockedBindings) Update(txid uint32, key, value string, opts *store.UpdateOptions) error {
	if key == "err_one" {
		return easyerr.Error("update err")
	}

	return nil
}

func (m *MockedBindings) Get(txid uint32, key string) (*entities.PlugKV, error) {

	if key == "err_one" {
		return nil, easyerr.Error("get err")
	}

	return &entities.PlugKV{
		Key:      key,
		Value:    "value1",
		Version:  int64(2),
		PlugsID:  "",
		TenantID: "",
	}, nil

}

func (m *MockedBindings) Del(txid uint32, key string) error {
	if key == "err_one" {
		return easyerr.Error("del err")
	}

	return nil
}

func (m *MockedBindings) DelBatch(txid uint32, keys []string) error {
	if keys[0] == "err_one" {
		return easyerr.Error("del batch err")
	}
	return nil
}

func (m *MockedBindings) Query(txid uint32, query *store.PkvQuery) ([]*entities.PlugKV, error) {
	if query.KeyPrefix == "err_one" {
		return nil, easyerr.Error("query err")
	}

	return []*entities.PlugKV{
		{
			Key:     "key11",
			Value:   "value11",
			Version: 2,
		},
	}, nil
}

func (m *MockedBindings) NewTxn() (uint32, error)    { return 0, nil }
func (m *MockedBindings) RollBack(txid uint32) error { return nil }
func (m *MockedBindings) Commit(txid uint32) error   { return nil }

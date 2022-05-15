package service

import "github.com/temphia/temphia/backend/server/btypes/store"

type Syncer interface {
	store.SyncDB
	GetInnerDriver() interface{}
}

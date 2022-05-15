package service

import "github.com/temphia/core/backend/server/btypes/store"

type Syncer interface {
	store.SyncDB
	GetInnerDriver() interface{}
}

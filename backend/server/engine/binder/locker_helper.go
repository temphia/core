package binder

import (
	"sync"

	"github.com/temphia/temphia/backend/server/btypes/easyerr"
)

type locker struct {
	locks map[string]struct{}
	llock sync.Mutex
}

func (l *locker) LockWatch(key string) error {
	return nil
}
func (l *locker) Lock(key string, expiry int) error {

	l.llock.Lock()
	defer l.llock.Unlock()
	_, ok := l.locks[key]
	if ok {
		return easyerr.Error("already locked")
	}

	return nil
}

func (l *locker) LockRenew(key string, expiry int) error {
	return nil
}

func (l *locker) UnLock(key string) error {
	l.llock.Lock()
	defer l.llock.Unlock()
	_, ok := l.locks[key]
	if !ok {
		return easyerr.Error("already unlocked")
	}

	delete(l.locks, key)

	return nil
}

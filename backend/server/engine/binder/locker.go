package binder

import (
	"sync"
)

var selfLocker = locker{
	locks: make(map[string]struct{}),
	llock: sync.Mutex{},
}

var resLocker = locker{
	locks: map[string]struct{}{},
	llock: sync.Mutex{},
}

func (b *Binder) selfLKey(key string) string {
	return b.namespace + b.plugId + key
}

func (b *Binder) SelfLockWait(key string) error {
	return selfLocker.LockWatch(b.selfLKey(key))

}
func (b *Binder) SelfLock(key string, expiry int) error {
	return selfLocker.Lock(b.selfLKey(key), expiry)
}

func (b *Binder) SelfLockRenew(key string, expiry int) error {
	return selfLocker.LockRenew(b.selfLKey(key), expiry)
}

func (b *Binder) SelfUnLock(key string) error {
	return selfLocker.UnLock(b.selfLKey(key))
}

func (b *Binder) resLKey(res, key string) string {
	return b.namespace + res + key
}

func (b *Binder) ResourceLockWait(resource string, key string) error {
	return resLocker.LockWatch(b.resLKey(resource, key))
}

func (b *Binder) ResourceLock(resource string, key string, expiry int) error {
	return resLocker.Lock(b.resLKey(resource, key), expiry)
}

func (b *Binder) ResourceLockRenew(resource string, key string, expiry int) error {
	return resLocker.LockRenew(b.resLKey(resource, key), expiry)
}

func (b *Binder) ResourceUnLock(resource string, key string) error {
	return resLocker.UnLock(b.resLKey(resource, key))
}

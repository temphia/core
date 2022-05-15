package binder

import (
	"sync"

	"github.com/temphia/temphia/backend/server/btypes/easyerr"
)

var _pcache = map[string][]byte{} // fixme => use proper cache libary
var _slock = sync.Mutex{}

func (b *Binder) Put(key string, value []byte, expire int) error {
	_slock.Lock()
	defer _slock.Unlock()

	_pcache[b.namespace+b.plugId+key] = value
	return nil
}
func (b *Binder) Get(key string) ([]byte, error) {
	_slock.Lock()
	defer _slock.Unlock()

	val, ok := _pcache[b.namespace+b.plugId+key]
	if !ok {
		return nil, easyerr.NotFound()
	}
	return val, nil
}

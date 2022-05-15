package runtime

import (
	"container/list"

	"github.com/temphia/temphia/backend/server/btypes/rtypes"
)

type simplePool struct {
	inner          map[string]*Cache
	maxUniquePlugs int
	maxBinders     int
}

func (p *simplePool) borrow(tenantId, plugId, agentId string) rtypes.ExecutorBinder {

	cache, ok := p.inner[tenantId]
	if !ok {
		return nil
	}

	binderCh, ok := cache.Get(Key(plugId + agentId))
	if !ok {
		return nil
	}
	select {
	case binder := <-binderCh:
		return binder
	default:
		return nil
	}
}

func (p *simplePool) returnn(tenantId, plugId, agentId string, binder rtypes.ExecutorBinder) {
	cache, ok := p.inner[tenantId]
	if !ok {
		cache = &Cache{
			ll:         list.New(),
			cache:      make(map[Key]*list.Element),
			MaxEntries: p.maxUniquePlugs,
			OnEvicted:  nil,
		}
		p.inner[tenantId] = cache
	}

	binderCh, ok := cache.Get(Key(plugId + agentId))
	if !ok {
		binderCh = make(chan rtypes.ExecutorBinder, p.maxBinders)
		cache.Add(Key(plugId), binderCh)
	}

	select {
	case binderCh <- binder:
	default:
	}
}

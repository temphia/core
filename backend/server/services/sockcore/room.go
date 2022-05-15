package sockcore

import (
	"strings"
	"sync"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/thoas/go-funk"
)

type RoomCore struct {
	ns          string
	name        string
	connections map[string][]string
	rlock       sync.Mutex
}

func (s *SockCore) getRoom(ns, name string, autoCreate bool) *RoomCore {

	var room *RoomCore
	key := ns + name

	s.roomLock.RLock()

	room = s.rooms[key]
	s.roomLock.RUnlock()

	if room != nil || !autoCreate {
		return room
	}

	s.roomLock.Lock()
	defer s.roomLock.Unlock()

	room = s.rooms[key]
	if room != nil {
		return room
	}

	room = &RoomCore{
		ns:          ns,
		name:        name,
		connections: make(map[string][]string),
		rlock:       sync.Mutex{},
	}
	s.rooms[key] = room

	return room
}

func (s *SockCore) roomClose(ns, roomId string) {
	s.roomLock.Lock()
	defer s.roomLock.Unlock()
	delete(s.rooms, ns+roomId)
}

func (r *RoomCore) addConn(connId string, _append bool, tags []string) {
	r.rlock.Lock()
	defer r.rlock.Unlock()

	pp.Println("@@ADD CONN ROOM", r.name, connId, tags)

	if _append {
		oldTags, ok := r.connections[connId]
		if ok {
			// fixme => maybe dedup

			tags = append(tags, oldTags...)
		}
	}

	r.connections[connId] = tags
}

func (r *RoomCore) removeConnTags(connId string, tags []string) {
	r.rlock.Lock()
	defer r.rlock.Unlock()

	oldTags, ok := r.connections[connId]
	if !ok {
		return
	}

	newTags := funk.LeftJoinString(oldTags, tags)

	if len(newTags) == 0 {
		delete(r.connections, connId)
		return
	}

	r.connections[connId] = newTags
}

func (r *RoomCore) updateTags(connId string, opts *service.UpdateTagOptions) error {
	r.rlock.Lock()
	defer r.rlock.Unlock()
	oldTags, ok := r.connections[connId]
	if !ok {
		return easyerr.NotFound()
	}
	newTags := make([]string, 0)

	// fixme => opts.AddTags forbid tags with sys. ?

	if opts.ClearOld {
		for _, v := range oldTags {
			if !strings.HasPrefix(v, "sys.") {
				continue
			}
			newTags = append(newTags, v)
		}
	} else {
		for _, v := range oldTags {
			if strings.HasPrefix(v, "sys.") {
				newTags = append(newTags, v)
				continue
			}
			if funk.ContainsString(opts.RemoveTags, v) {
				continue
			}
			newTags = append(newTags, v)
		}
	}

	newTags = append(newTags, opts.AddTags...)
	r.connections[connId] = newTags
	return nil
}

func (s *RoomCore) getAllConns() []string {

	s.rlock.Lock()
	defer s.rlock.Unlock()

	keys := make([]string, 0, len(s.connections))
	for k := range s.connections {
		keys = append(keys, k)
	}

	return keys
}

func (s *RoomCore) getAllConnsWithTags() map[string][]string {
	s.rlock.Lock()
	defer s.rlock.Unlock()

	allConns := make(map[string][]string, len(s.connections))
	for k, v := range s.connections {
		allConns[k] = v
	}

	return allConns
}

func (s *RoomCore) getTagConns(tags []string) []string {
	s.rlock.Lock()
	defer s.rlock.Unlock()

	connKeys := make([]string, 0, len(s.connections))
	for connId, connTags := range s.connections {

		skip := false
		for _, tag := range tags {
			if !funk.ContainsString(connTags, tag) {
				skip = true
			}
		}
		if skip {
			continue
		}
		connKeys = append(connKeys, connId)
	}

	return connKeys
}

func (s *RoomCore) getConns(conns []string) []string {
	s.rlock.Lock()
	defer s.rlock.Unlock()

	keys := make([]string, 0, len(s.connections))
	for k := range s.connections {
		if funk.ContainsString(conns, k) {
			keys = append(keys, k)
		}
	}

	return keys
}

func (r *RoomCore) removeConn(connId string) {
	r.rlock.Lock()
	defer r.rlock.Unlock()
	delete(r.connections, connId)

}

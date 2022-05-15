package sockcore

import (
	"log"
	"time"

	"github.com/k0kubun/pp"
	"github.com/thoas/go-funk"
)

func (s *SockCore) writeConns(ns string, data []byte, conns []string, ignores ...string) {

	ignore := len(ignores) == 0

	for _, c := range conns {

		conn, ok := s.connections[ns+c]
		if !ok {
			pp.Println("SKIPPING", ns, c)
			continue
		}

		if ignore && funk.ContainsString(ignores, c) {
			pp.Println("IGNOREING", ns, c)
			continue
		}

		pp.Println("WRITING", c)

		conn.write(data)
	}
}

func (s *SockCore) garbageCollect() {
	time.Sleep(1 * time.Hour)

	for {
		shortlist := make([]string, 0)

		current := time.Now().UnixNano()

		for k, c := range s.connections {
			if c.expiry < int(current) {
				shortlist = append(shortlist, k)
			}
		}

		for _, connKey := range shortlist {
			func() {
				s.connLock.Lock()
				defer s.connLock.Unlock()

				conn, ok := s.connections[connKey]
				if !ok {
					return
				}
				if conn.expiry < int(current) {
					delete(s.connections, connKey)
					conn.close()
				}
			}()
		}

		time.Sleep(30 * time.Minute)
	}

}

func (s *SockCore) _broadcastFromClient(ns, room, fromId string, data []byte) {
	_room := s.getRoom(ns, room, false)
	connIds := _room.getAllConns()

	for idx, connId := range connIds {
		if connId == fromId {
			connIds[idx] = ""
		}
	}

	s.writeConns(ns, data, connIds)

	if s.peerRouter != nil {
		err := s.peerRouter.RouteClientBroadcast(ns, room, fromId, data)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *SockCore) _ticketMsgFromClient(ns, room, fromId, ticket string, data []byte, targetIds ...string) {
	pp.Println("@ticket", fromId)
	s.writeConns(ns, data, targetIds)

	if s.peerRouter != nil {
		err := s.peerRouter.RouteClientPublishTicket(ns, room, fromId, ticket, data)
		if err != nil {
			log.Println(err)
		}
	}

}

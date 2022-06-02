package sockcore

import (
	"sync"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/service"
)

var _ service.SockCore = (*SockCore)(nil)

type SockCore struct {
	rooms    map[string]*RoomCore
	roomLock sync.RWMutex
	name     string

	connections map[string]*Conn
	connLock    sync.Mutex
	ticketParse TicketParse
	peerRouter  service.SockPeerRouter
	stateSyncer service.SockStateSyncer
}

func New() *SockCore {
	s := &SockCore{
		rooms:       make(map[string]*RoomCore),
		connections: make(map[string]*Conn),
		roomLock:    sync.RWMutex{},
		connLock:    sync.Mutex{},
	}
	go s.debug()
	return s
}

func (s *SockCore) NewConnection(opts *service.ConnOptions) error {
	s.connLock.Lock()
	defer s.connLock.Unlock()
	_conn := &Conn{
		parent:  s,
		conn:    opts.Conn,
		ns:      opts.NameSpace,
		closed:  false,
		expiry:  opts.Expiry,
		writeCh: make(chan []byte),
	}

	connKey := opts.NameSpace + opts.Conn.Id()
	old := s.connections[connKey]
	if old != nil {
		go func() {
			old.close()
		}()
	}

	_conn.start()

	s.connections[connKey] = _conn

	if opts.PreJoinRooms == nil {
		return nil
	}

	for k, v := range opts.PreJoinRooms {
		room := s.getRoom(opts.NameSpace, k, true)
		room.addConn(opts.Conn.Id(), false, v)
		// fixme => actually publish to "sockd state store thing"
	}

	pp.Println("PRE_CONNECT_ROOMS", opts)
	pp.Println(s.rooms)

	return nil
}

func (s *SockCore) SendBroadcast(ns, room string, ignoreConns []string, payload []byte) error {
	msg := &message{
		Room:        room,
		Type:        MESSAGE_SERVER_BROADCAST,
		XID:         xid.New().String(),
		Payload:     payload,
		ServerIdent: s.name,
		Ticket:      "",
		FromId:      "",
		Targets:     nil,
	}

	bytes, err := msg.JSON()
	if err != nil {
		return err
	}

	room5 := s.getRoom(ns, room, false)
	if room5 == nil {
		return easyerr.Error("Room not found")
	}

	pp.Println(room5, ns, room)

	conns := room5.getAllConns()

	pp.Println("BROADCASTING TO ", len(conns))

	s.writeConns(ns, bytes, conns, ignoreConns...)

	return nil
}

func (s *SockCore) SendTagged(ns, room string, tags []string, ignoreConns []string, payload []byte) error {
	pp.Println(room, tags, ignoreConns)

	msg := &message{
		Room:        room,
		Type:        MESSAGE_SERVER_PUBLISH,
		XID:         xid.New().String(),
		Payload:     payload,
		ServerIdent: s.name,
		FromId:      "",
		Ticket:      "",
		Targets:     nil,
	}

	bytes, err := msg.JSON()
	if err != nil {
		return err
	}

	room5 := s.getRoom(ns, room, false)
	if room5 == nil {

		pp.Println(s.rooms)

		return easyerr.NotFound()
	}

	conns := room5.getTagConns(tags)

	s.writeConns(ns, bytes, conns, ignoreConns...)

	return nil
}

func (s *SockCore) SendDirect(ns, room string, connId []string, payload []byte) error {
	msg := &message{
		Type:        MESSAGE_SERVER_DIRECT,
		XID:         xid.New().String(),
		Payload:     payload,
		Room:        "",
		FromId:      "",
		ServerIdent: "",
		Ticket:      "",
		Targets:     nil,
	}
	bytes, err := msg.JSON()
	if err != nil {
		return err
	}

	s.writeConns(ns, bytes, connId)
	return nil
}

// meta methods
func (s *SockCore) RoomAddConn(ns, roomId, connId string, append bool, tags []string) error {
	_room := s.getRoom(ns, roomId, true)
	_room.addConn(connId, append, tags)
	return nil
}

func (s *SockCore) RoomRemoveTags(ns, roomId, connId string, tags []string) error {
	_room := s.getRoom(ns, roomId, false)
	if _room == nil {
		return easyerr.NotFound()
	}
	_room.removeConnTags(connId, tags)
	return nil
}

func (s *SockCore) RoomUpdateTags(ns, roomId, connId string, opts *service.UpdateTagOptions) error {

	_room := s.getRoom(ns, roomId, false)
	if _room == nil {
		return easyerr.NotFound()
	}

	return _room.updateTags(connId, opts)
}

func (s *SockCore) RoomListConn(ns, roomId string) (map[string][]string, error) {
	_room := s.getRoom(ns, roomId, false)
	if _room == nil {
		return nil, easyerr.NotFound()
	}

	return _room.getAllConnsWithTags(), nil

}

func (s *SockCore) RoomClose(ns, roomId string) error {
	s.roomClose(ns, roomId)
	return nil
}

func (s *SockCore) RoomKickConn(ns, roomId, connId string) error {
	_room := s.getRoom(ns, roomId, false)
	if _room == nil {
		return easyerr.NotFound()
	}
	_room.removeConn(connId)
	return nil
}

func (s *SockCore) BannConn(ns, connId string) error {
	s.connLock.Lock()
	defer s.connLock.Unlock()
	delete(s.connections, ns+connId)
	return nil
}

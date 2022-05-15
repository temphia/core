package service

type Envelope struct {
	// sock peer message
	// encode with cobor

	OriginServer string
	InnerMessage []byte
	TargetType   string
	Targets      []string
	Room         string
}

type Conn interface {
	Id() string
	Write([]byte) error
	Close() error
	Read() ([]byte, error)
}

type ConnOptions struct {
	NameSpace    string
	Conn         Conn
	Expiry       int
	PreJoinRooms map[string][]string
}

type UpdateTagOptions struct {
	AddTags    []string
	ClearOld   bool
	RemoveTags []string
}

type SockCore interface {
	NewConnection(opts *ConnOptions) error
	SendDirect(ns, room string, connId []string, payload []byte) error
	SendBroadcast(ns, room string, ignoreConns []string, payload []byte) error
	SendTagged(ns, room string, tags []string, ignoreConns []string, payload []byte) error

	RoomAddConn(ns, roomId, connId string, append bool, tags []string) error
	RoomRemoveTags(ns, roomId, connId string, tags []string) error
	RoomUpdateTags(ns, roomId, connId string, opts *UpdateTagOptions) error

	RoomListConn(ns, roomId string) (map[string][]string, error)
	RoomClose(ns, roomId string) error
	RoomKickConn(ns, roomId, connId string) error
	BannConn(ns, connId string) error
}

type SockPeerRouter interface {
	RouteBroadcast(ns, room string, payload []byte) error
	RoutePublishTagged(ns, room string, payload []byte, tags []string) error
	RouteSendConn(ns, connId string, payload []byte) error
	RouteClientBroadcast(ns, room, clientId string, data []byte) error
	RouteClientPublishTicket(ns, room, fromId, ticket string, data []byte) error
}

type SockStateSyncer interface {
	SyncRoomAddConn(ns, roomId, connId string, append bool, tags []string) error
	SyncRoomRemoveTags(ns, roomId, connId string, tags []string) error
	SyncRoomListConn(ns, roomId string) ([]string, error)
	SyncRoomClose(ns, roomId string) error
	SyncRoomKickConn(ns, roomId, connId string) error
	SyncBannConn(ns, connId string) error
}

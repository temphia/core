package binder

import "github.com/temphia/core/backend/server/btypes/service"

func (b *Binder) sockd() service.SockCore {
	return b.factory.Sockd
}

func (b *Binder) SendDirect(room string, connId []string, payload []byte) error {
	return b.sockd().SendDirect(b.namespace, b.roomResource(room), connId, payload)
}

func (b *Binder) SendBroadcast(room string, payload []byte) error {
	return b.sockd().SendBroadcast(b.namespace, b.roomResource(room), []string{}, payload)
}

func (b *Binder) SendTagged(room string, tags []string, ignoreConns []string, payload []byte) error {
	return b.sockd().SendTagged(b.namespace, b.roomResource(room), tags, ignoreConns, payload)
}

func (b *Binder) AddToRoom(room string, connId string, tags []string) error {
	return b.sockd().RoomAddConn(b.namespace, b.roomResource(room), connId, true, tags)
}

func (b *Binder) KickFromRoom(room string, connId string) error {
	return b.sockd().RoomKickConn(b.namespace, b.roomResource(room), connId)
}

func (b *Binder) ListRoomConns(room string) (map[string][]string, error) {
	return b.sockd().RoomListConn(b.namespace, b.roomResource(room))
}

func (b *Binder) BannConn(connId string) error {
	return b.sockd().BannConn(b.namespace, connId)
}

// private

func (b *Binder) roomResource(name string) string {
	// if name == "plugs_dev" && b.binder.Plug.Dev {
	// 	return name
	// }

	return b.resourced(name)
}

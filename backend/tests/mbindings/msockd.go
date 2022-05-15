package mbindings

import "github.com/temphia/core/backend/server/btypes/easyerr"

func (m *MockedBindings) SendDirect(room string, connId []string, payload []byte) error {

	if "err_room" == room {
		return easyerr.Error("senddirect sockd err")
	}

	return nil
}

func (m *MockedBindings) SendBroadcast(room string, payload []byte) error {
	if "err_room" == room {
		return easyerr.Error("sendbroadcast sockd err")
	}

	return nil
}

func (m *MockedBindings) SendTagged(room string, tags []string, ignoreConns []string, payload []byte) error {
	if "err_room" == room {
		return easyerr.Error("sendtagged sockd err")
	}

	return nil
}

func (m *MockedBindings) AddToRoom(room string, connId string, tags []string) error {
	if "err_room" == room {
		return easyerr.Error("addtoroom sockd err")
	}

	return nil
}

func (m *MockedBindings) KickFromRoom(room string, connId string) error {
	if "err_room" == room {
		return easyerr.Error("kickfromroom sockd err")
	}

	return nil
}

func (m *MockedBindings) ListRoomConns(room string) (map[string][]string, error) {
	if "err_room" == room {
		return nil, easyerr.Error("listroomconns sockd err")
	}

	return nil, nil
}

func (m *MockedBindings) BannConn(connId string) error {
	return nil
}

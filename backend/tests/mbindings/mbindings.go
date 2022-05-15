package mbindings

import (
	"log"
	"time"

	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
)

// var _ rtypes.Bindings = (*MockedBindings)(nil)

// this does not test functional logic of the bindings
// this only checks if it passes arguments and return types
// correctly

type MockedBindings struct {
	SelfFiles map[string][]byte
}

type Value struct {
	Inner    string
	Version  int
	Tag      string
	Audience string
	Source   string
	TTL      int64
}

func (m *MockedBindings) Log(msg string) {
	log.Println(msg)
}

func (m *MockedBindings) LazyLog(msgs []string) {
	log.Println(msgs)
}

func (m *MockedBindings) Sleep(t int) {
	time.Sleep(time.Second * time.Duration(t))
}

func (m *MockedBindings) GetSelfFile(file string) ([]byte, error) {
	data, ok := m.SelfFiles[file]

	if !ok {
		return nil, easyerr.Error("not found")
	}
	return data, nil
}

func (m *MockedBindings) GetPlugKVBindings() rtypes.PlugKVBindings   { return nil }
func (m *MockedBindings) GetSockdBindings() rtypes.SockdBindings     { return nil }
func (m *MockedBindings) GetUserBindings() rtypes.UserBindings       { return nil }
func (m *MockedBindings) GetCabinetBindings() rtypes.CabinetBindings { return nil }
func (m *MockedBindings) GetSelfBindings() rtypes.SelfBindings       { return nil }

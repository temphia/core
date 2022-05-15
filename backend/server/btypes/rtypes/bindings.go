package rtypes

import (
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/store"
)

type Bindings interface {
	BindingsCore

	GetPlugKVBindings() PlugKVBindings
	GetSockdBindings() SockdBindings
	GetUserBindings() UserBindings
	GetCabinetBindings() CabinetBindings
	GetSelfBindings() SelfBindings
	GetPCache() PCache
}

type BindingsCore interface {
	Log(msg string)
	LazyLog(msgs []string)
	Sleep(int)
	GetSelfFile(file string) ([]byte, error)

	ContextBindings
	GetApp() interface{}
}

type ContextBindings interface {
	GetRequestData() interface{}
	GetRequestVar(vname string) interface{}
	GetRequestVars() map[string]interface{}
	SetResponse(interface{})
	SetResponseVar(name string, value interface{})
}

type PlugKVBindings interface {
	Set(txid uint32, key, value string, opts *store.SetOptions) error
	Update(txid uint32, key, value string, opts *store.UpdateOptions) error
	Get(txid uint32, key string) (*entities.PlugKV, error)
	Del(txid uint32, key string) error
	DelBatch(txid uint32, keys []string) error
	Query(txid uint32, query *store.PkvQuery) ([]*entities.PlugKV, error)

	NewTxn() (uint32, error)
	RollBack(txid uint32) error
	Commit(txid uint32) error
}

type CabinetBindings interface {
	AddFile(bucket string, file string, contents []byte) error
	ListFolder(bucket string) ([]string, error)
	GetFile(bucket string, file string) ([]byte, error)
	DeleteFile(bucket string, file string) error
	GenerateTicket(bucket string, ticket *CabTicket) (string, error)
}

type SockdBindings interface {
	SendDirect(room string, connId []string, payload []byte) error
	SendBroadcast(room string, payload []byte) error
	SendTagged(room string, tags []string, ignoreConns []string, payload []byte) error

	AddToRoom(room string, connId string, tags []string) error
	KickFromRoom(room string, connId string) error
	ListRoomConns(room string) (map[string][]string, error)
	BannConn(connId string) error
}

type UserBindings interface {
	ListUsers(group string) ([]string, error)
	MessageUser(group, user, message string, encrypted bool) error
	GetUser(group, user string) (*entities.UserInfo, error)

	MessageCurrentUser(user, message string, encrypted bool) error
	CurrentUser() (*entities.UserInfo, error)
}

type LockerBindinigs interface {
	// fixme => nested key lock

	SelfLockWait(key string) error
	SelfLock(key string, expiry int) error
	SelfLockRenew(key string, expiry int) error
	SelfUnLock(key string) error

	ResourceLockWait(resource string, key string) error
	ResourceLock(resource string, key string) error
	ResourceLockRenew(resource string, key string) error
	ResourceUnLock(resource string, key string) error
}

type NetBindings interface {
	HTTPCacheFile(url string, headers map[string]string) error
	HTTPCall(HTTPRequest) *HTTPResponse
	QuickJsonGet(url string) ([]byte, error)
	QuickJsonPost(url string, data []byte) ([]byte, error)
}

type PCache interface {
	Put(key string, value []byte, expire int) error
	Get(key string) ([]byte, error)
}

type SelfBindings interface {
	SelfGetFile(file string) ([]byte, error)
	SelfAddFile(file string, data []byte) error
	SelfUpdateFile(file string, data []byte) error
	SelfListResources() ([]*Resource, error)
	SelfGetResource(name string) (*Resource, error)
	SelfIncomingConns() ([]*Connection, error)
	SelfOutgoingConns() ([]*Connection, error)

	SelfExecuteSlot(name string, opts SlotOption) ([]byte, error)
	SelfRenderFile(file string, opts RenderOption) ([]byte, error)
}

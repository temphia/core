package btypes

import (
	"github.com/rs/zerolog"
	"github.com/temphia/core/backend/server/btypes/store"

	"gitlab.com/mr_balloon/golib/hmap"
)

const (
	DefaultTenant     = "default0"
	BprintBlobFolder  = "bprints"
	BPrintTypePlug    = "plug"
	BprintTypePlugin  = "plugin"
	BPrintTypeTSchema = "tschema"
)

const (
	CORE_DB_VER_MIN      = 0
	CORE_DB_VER_MAX      = 0
	PLUGSTATE_VER_MIN    = 0
	PLUGSTATE_VER_MAX    = 0
	EVENT_STORE_VER_MIN  = 0
	EVENT_STORE_VER_MAX  = 0
	DTABLE_STORE_VER_MIN = 0
	DTABLE_STORE_VER_MAX = 0
)

type Data interface {
	AsJsonBytes() ([]byte, error)
	AsStruct(interface{}) error
	AsHMap() (hmap.H, error)
}

type Server interface {
	BindRoutes()
	Listen() error
	Close() error
}

type App interface {
	NodeId() string
	DevMode() bool

	Registry() interface{}
	Run() error
	ControlPlane() interface{}
	Fencer() interface{}
	Engine() interface{}
	Sockd() interface{}
	CoreHub() store.CoreHub
	PlugKV() store.PlugStateKV
	Signer() interface{}
	Cabinet() store.CabinetHub
	Pacman() interface{}
	DynHub() store.DynHub
	Data() store.AssetStore
	Server() Server
	RootController() interface{}
	RootLogger() *zerolog.Logger
	SingleTenant() bool
	TenantId() string
}

var DevMode = true

const ROOM_SYSTABLE = "sys.dtable"
const ROOM_SYS_USERS = "sys.users"
const ROOM_PLUG_DEV = "plugs_dev"

const TAG_REALUSER = "sys.real_user"
const TAG_CONSOLE_CONN = "sys.console_conn"

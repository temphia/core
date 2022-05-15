package env

import (
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/models/entities"
)

type EnvType string

const (
	SignalEnv        EnvType = "signal"
	EventEnv         EnvType = "event"
	ResourceEnv      EnvType = "resource"
	ConsoleInitEnv   EnvType = "console_init"
	ConsoleActionEnv EnvType = "console_action"
)

type Signal struct {
	Type        EnvType
	Inner       *entities.Resource
	ToPlug      *entities.Plug
	ToAgent     *entities.Agent
	ToResources map[string]*entities.Resource
	Async       bool
	ExecDepth   uint32
	Parent      string
	Payload     btypes.Data
}

type ResourceEval struct {
	Type   EnvType
	Inner  *entities.Resource
	Action string
}

type Module struct {
	Type    EnvType
	Name    string
	Path    string
	Method  string
	Payload btypes.Data
}

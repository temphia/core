package vmodels

import (
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
)

type ExecutorData struct {
	Plug      *entities.Plug
	Agent     *entities.Agent
	Bprint    *entities.BPrint
	Resources map[string]*entities.Resource
}

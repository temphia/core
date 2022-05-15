package plugloader

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/models/vmodels"
	"github.com/temphia/core/backend/server/btypes/store"
)

func Load(syncer store.SyncDB, tenantId, plugId, agentId string) (*vmodels.ExecutorData, error) {
	plug, err := syncer.PlugGet(tenantId, plugId)
	if err != nil {
		pp.Println("err@plug")
		return nil, err
	}

	agent, err := syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		pp.Println("err@agent")

		return nil, err
	}

	pp.Println("AGENT & PLUG loaded")

	// bprint err is ignore so bprint might be nil
	// so we have to check that in binder
	bprint, _ := syncer.BprintGet(tenantId, plug.BprintId)
	resIds := make([]string, 0, len(agent.Resources))
	for _, v := range agent.Resources {
		resIds = append(resIds, v)
	}

	resources, err := syncer.ResourcesMulti(tenantId, resIds...)
	if err != nil {
		pp.Println("err@resources")
		return nil, err
	}

	fresources := make(map[string]*entities.Resource)
	for k, v := range agent.Resources {
		for _, respRes := range resources {
			if respRes.Id != v {
				continue
			}
			fresources[k] = respRes
		}
	}

	return &vmodels.ExecutorData{
		Plug:      plug,
		Agent:     agent,
		Bprint:    bprint,
		Resources: fresources,
	}, nil

}

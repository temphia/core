package hoster

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/rtypes/job"
	"github.com/temphia/core/backend/server/btypes/store"
)

func (m *Manager) processTSP(csource store.CabinetSourced, c *gin.Context, folder, file string) {

	runtime := m.engine.GetRuntime()

	// fixme => move this to engine

	j := &job.Job{
		PlugId:  "",
		AgentId: "",
		EventId: "",
	}
	runtime.Shedule(j)
	j.Wait()

	resp, err := j.Result()
	if err != nil {
		pp.Println(err)
		return
	}

	pp.Println(resp)

}

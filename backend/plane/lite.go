package plane

import (
	"github.com/temphia/core/backend/server/app/config"
	"github.com/temphia/core/backend/server/btypes/rtypes/job"
	"github.com/temphia/core/backend/server/btypes/service"
)

var _ service.ControlPlane = (*PlaneLite)(nil)

type PlaneLite struct {
	eventbus *EventBus
	locker   *Locker
	router   *Router
}

func NewLite() *PlaneLite {

	return &PlaneLite{
		eventbus: NewEventBus(),
		locker:   NewLocker(),
		router:   nil,
	}
}

// dl and start up stuff

func (p *PlaneLite) Start() error                                            { return nil }
func (p *PlaneLite) Inject(iapp interface{}, config *config.AppConfig) error { return nil }

// liveness and status stuff

func (p *PlaneLite) NotifyStat(stats service.NodeStats) error    { return nil }
func (p *PlaneLite) GetAppStats() (*service.ClusterStats, error) { return nil, nil }

// router stuff
func (p *PlaneLite) SetJobChan(ic chan *job.Job) {
	p.router = NewRouter(ic)
}

func (p *PlaneLite) GetRouter() service.Router {
	return p.router
}

// locker
func (p *PlaneLite) GetLocker() service.Locker {
	return nil
}

// sockdrouter
func (p *PlaneLite) GetSockdRouter() service.SockdRouter {
	return nil
}

// eventbus
func (p *PlaneLite) GetEventBus() service.EventBus { return p.eventbus }

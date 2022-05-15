package plane

import (
	"github.com/temphia/core/backend/server/btypes/rtypes/job"
)

type Router struct {
	// natsConn *nats.Conn
	inJob chan *job.Job
}

func NewRouter(inchan chan *job.Job) *Router {
	return &Router{
		inJob: inchan,
	}
}

func (r Router) Route(j *job.Job) bool {
	// fixme => route to peers here ?

	return false
}

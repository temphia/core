package rtypes

import (
	"github.com/temphia/temphia/backend/server/btypes/env"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/job"
	"github.com/temphia/temphia/backend/server/lib"
)

type Runtime interface {
	Run() error
	Shedule(j *job.Job)
}

type RuntimeBinder interface {
	GetModule(string) (Module, error)
	Signal(*env.Signal) (lib.LazyMap, error)
}

type Router interface {
	Route(j *job.Job) bool
}

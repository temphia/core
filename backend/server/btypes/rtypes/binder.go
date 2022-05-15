package rtypes

import (
	"github.com/temphia/core/backend/server/btypes/rtypes/job"
)

type ExecutorBinder interface {
	Bindings

	AttachJob(j *job.Job)
	Execute()
}

// SysPlug

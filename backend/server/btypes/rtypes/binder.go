package rtypes

import (
	"github.com/temphia/temphia/backend/server/btypes/rtypes/job"
)

type ExecutorBinder interface {
	Bindings

	AttachJob(j *job.Job)
	Execute()
}

// SysPlug

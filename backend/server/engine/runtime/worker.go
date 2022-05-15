package runtime

import (
	"log"
	"time"

	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/rtypes/job"
	"github.com/ztrue/tracerr"
)

func (r *runtime) shedule(j *job.Job) {

	// fixme => check policy here
	if r.router != nil {
		if r.router.Route(j) {
			return
		}
	}

	r.jobCh <- j
}

func (r *runtime) counterPlus(b rtypes.ExecutorBinder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()

	r.runningExecs[b] = struct{}{}
}

func (r *runtime) counterMinus(b rtypes.ExecutorBinder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()

	delete(r.runningExecs, b)
}

func (r *runtime) worker() {

	for {
		r.doWork(<-r.jobCh)
	}

}

func (r *runtime) doWork(j *job.Job) {
	binder, err := r.getBinder(j)
	if err != nil {
		j.Err(tracerr.Wrap(err))
		return
	}

	r.counterPlus(binder)
	binder.Execute()
	r.counterMinus(binder)
	r.setBinder(j, binder)
}

func (r *runtime) debugLoop() {

	for {
		time.Sleep(time.Second * 5)

		currentExecs := make(map[rtypes.ExecutorBinder]struct{})
		for k, v := range r.runningExecs {
			currentExecs[k] = v
		}

		for exec := range currentExecs {
			// info := exec.GetInfo()
			log.Println(exec)
		}
	}
}

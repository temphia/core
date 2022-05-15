package runtime

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/rtypes/job"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/temphia/core/backend/server/engine/binder"
	"github.com/temphia/core/backend/server/registry"
)

var (
	_ rtypes.Runtime = (*runtime)(nil)
)

type runtime struct {
	close        chan struct{}
	jobCh        chan *job.Job
	runningExecs map[rtypes.ExecutorBinder]struct{}
	mlock        sync.Mutex

	router      rtypes.Router
	localFolder string
	blobfolder  string
	pool        simplePool

	app           btypes.App
	execBuilders  map[string]rtypes.ExecutorBuilder
	binderFactory binder.Factory

	//ext services
	fencer service.Fencer
	signer service.Signer
	syncer store.SyncDB
	logger zerolog.Logger
}

func New(_app btypes.App, logger zerolog.Logger) *runtime {
	rt := &runtime{
		close:        make(chan struct{}),
		jobCh:        make(chan *job.Job),
		runningExecs: make(map[rtypes.ExecutorBinder]struct{}),
		mlock:        sync.Mutex{},
		router:       nil,
		app:          _app,
		execBuilders: nil, // fixme
		fencer:       _app.Fencer().(service.Fencer),
		signer:       _app.Signer().(service.Signer),
		syncer:       _app.CoreHub(),
		logger:       logger,
		localFolder:  "",
		blobfolder:   "",
		pool: simplePool{
			inner:          make(map[string]*Cache),
			maxUniquePlugs: 100,
			maxBinders:     100,
		},
	}

	return rt
}

func (r *runtime) Run() error {
	reg := r.app.Registry().(*registry.Registry)
	r.execBuilders = reg.GetExecutors()
	r.binderFactory = binder.Factory{
		App:    r.app,
		Sockd:  r.app.Sockd().(service.SockCore),
		Pacman: r.app.Pacman().(service.Pacman),
		Logger: r.logger,
	}

	go r.worker()
	return nil
}

func (r *runtime) Shedule(j *job.Job) {
	r.shedule(j)
}

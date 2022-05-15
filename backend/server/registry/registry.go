package registry

import (
	"errors"
	"sync"

	"github.com/temphia/temphia/backend/server/app/config"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/store"
)

type (
	StoreBuilder func(*config.StoreSource) (store.Store, error)

	ExecModuleBuilder func(rtypes.ModuleOption) (rtypes.Module, error)
	DynamicScript     func(ns string, ctx interface{}) error
)

type Registry struct {
	repoBuilders   map[string]store.RepoBuilder
	executors      map[string]rtypes.ExecutorBuilder
	execModules    map[string]ExecModuleBuilder
	dynamicScripts map[string]DynamicScript
	storeBuilders  map[string]StoreBuilder
	freezed        bool
	mlock          *sync.Mutex
}

var (
	errTooLate = errors.New("err too late")
	errTooSoon = errors.New("err too soon")
)

func New(fromGlobal bool) *Registry {
	reg := &Registry{
		freezed:        false,
		dynamicScripts: make(map[string]DynamicScript),
		repoBuilders:   make(map[string]store.RepoBuilder),
		executors:      make(map[string]rtypes.ExecutorBuilder),
		execModules:    make(map[string]ExecModuleBuilder),
		storeBuilders:  make(map[string]StoreBuilder),
		mlock:          &sync.Mutex{},
	}

	if !fromGlobal || G == nil {
		return reg
	}

	G.mlock.Lock()
	defer G.mlock.Unlock()

	for k, v := range G.storeBuilders {
		reg.storeBuilders[k] = v
	}

	for k, v := range G.repoBuilders {
		reg.repoBuilders[k] = v
	}

	for k, v := range G.executors {
		reg.executors[k] = v
	}
	for k, v := range G.execModules {
		reg.execModules[k] = v
	}
	for k, v := range G.dynamicScripts {
		reg.dynamicScripts[k] = v
	}

	return reg
}

func (r *Registry) Freeze() {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	r.freezed = true
}

func (r *Registry) SetRepoBuilder(name string, builder store.RepoBuilder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}
	r.repoBuilders[name] = builder
}

func (r *Registry) SetExecutor(name string, builder rtypes.ExecutorBuilder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}
	r.executors[name] = builder
}

func (r *Registry) SetExecModule(name string, builder ExecModuleBuilder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}
	r.execModules[name] = builder
}

func (r *Registry) SetDynamicScript(name string, script func(ns string, ctx interface{}) error) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}
	r.dynamicScripts[name] = script
}

func (r *Registry) SetStoreBuilder(name string, b StoreBuilder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}

	r.storeBuilders[name] = b
}

func (r *Registry) GetRepoBuilders() map[string]store.RepoBuilder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}
	return r.repoBuilders
}

func (r *Registry) GetExecutors() map[string]rtypes.ExecutorBuilder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}
	return r.executors
}

func (r *Registry) GetExecModules() map[string]ExecModuleBuilder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}
	return r.execModules
}

func (r *Registry) GetDynamicScripts() map[string]DynamicScript {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}

	return r.dynamicScripts
}

func (r *Registry) GetStoreBuilders() map[string]StoreBuilder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}

	return r.storeBuilders
}

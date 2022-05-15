package registry

import (
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/store"
)

var G *Registry

func init() {
	if G == nil {
		G = New(false)
	}

}

func SetRepoBuilder(name string, builder store.RepoBuilder) {
	G.SetRepoBuilder(name, builder)

}

func SetStoreBuilders(name string, builder StoreBuilder) {
	G.SetStoreBuilder(name, builder)
}

func SetExecutor(name string, builder rtypes.ExecutorBuilder) {
	G.SetExecutor(name, builder)
}

func SetExecModule(name string, builder ExecModuleBuilder) {
	G.SetExecModule(name, builder)
}

func SetDynamicScript(name string, script func(ns string, ctx interface{}) error) {
	G.SetDynamicScript(name, script)
}

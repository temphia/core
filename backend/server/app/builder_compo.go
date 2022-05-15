package app

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/temphia/core/backend/server/engine"
	"github.com/temphia/core/backend/server/services/cabinethub"
	"github.com/temphia/core/backend/server/services/corehub"
	"github.com/temphia/core/backend/server/services/dynhub"
	"github.com/temphia/core/backend/server/services/fencer"
	"github.com/temphia/core/backend/server/services/pacman"
	"github.com/temphia/core/backend/server/services/signer"
	"github.com/temphia/core/backend/server/services/sockcore"
)

func (b *Builder) buildComponents() error {
	b.app.signer = signer.New([]byte(b.config.MasterKey), "temphia")
	b.app.sockd = sockcore.New()

	err := b.buildStores()
	if err != nil {
		return err
	}
	err = b.buildPacman()
	if err != nil {
		return err
	}
	err = b.buildServices()
	if err != nil {
		return err
	}

	return nil

}

func (b *Builder) buildStores() error {
	b.registery.Freeze()

	storeBuilders := b.registery.GetStoreBuilders()

	for _, ss := range b.config.StoreSources {
		sBuilder := storeBuilders[ss.Provider]
		if sBuilder == nil {
			pp.Println(ss)
			fmt.Println(storeBuilders)

			return easyerr.Error(fmt.Sprintf("Provider %s not found", ss.Provider))
		}

		store, err := sBuilder(ss)
		if err != nil {
			return easyerr.Wrap("err while building store", err)
		}
		b.stores[ss.Name] = store
	}

	if _, ok := b.stores[b.config.StoreOptions.DefaultCabinet]; !ok {
		return easyerr.Error(fmt.Sprintf("default cabinet not found %s", b.config.StoreOptions.DefaultCabinet))
	}

	if _, ok := b.stores[b.config.StoreOptions.DefaultCoreDB]; !ok {
		return easyerr.Error(fmt.Sprintf("default coredb not found %s", b.config.StoreOptions.DefaultCoreDB))
	}

	if sp, ok := b.stores[b.config.StoreOptions.DefaultCoreDB]; !ok || !sp.Supports(store.TypeCoreDB) {
		return easyerr.Error(fmt.Sprintf("default coredb not loaded %s", b.config.StoreOptions.DefaultCoreDB))
	} else {
		b.coreDB = sp.CoreDB()

		b.app.plugKV = sp.StateDB()
	}

	csources := make(map[string]store.CabinetSource)
	for storek, _store := range b.stores {
		if !_store.Supports(store.TypeBlobStore) {
			continue
		}
		csources[storek] = _store.CabinetSource()
	}

	_, ok := csources[b.config.StoreOptions.DefaultCabinet]
	if !ok {
		return easyerr.Error(fmt.Sprintf("default cabinet not loaded %s", b.config.StoreOptions.DefaultCabinet))
	}

	fmt.Println("@=>", b.stores)

	b.app.cabinetHub = cabinethub.New(csources, b.config.StoreOptions.DefaultCabinet)

	return nil
}

func (b *Builder) buildPacman() error {
	repoBuilders := b.registery.GetRepoBuilders()

	for _, ss := range b.config.BprintSources {
		rBuilder := repoBuilders[ss.Provider]
		if rBuilder == nil {
			return easyerr.Error(fmt.Sprintf("Repo Provider not found: %s", ss.Provider))
		}

		repo, err := rBuilder(&store.RepoBuilderOptions{
			TenantId:      "",
			BaseURL:       "",
			SourceOptions: ss,
		})

		if err != nil {
			return err
		}

		b.repos[ss.Name] = repo
	}

	return nil

}

func (b *Builder) buildServices() error {
	b.app.registry = b.registery
	b.app.engine = engine.New(b.app, *b.rtLogger)
	b.app.fencer = fencer.New(b.app)

	{
		dyns := make(map[string]store.DynDB)
		for k, s := range b.stores {
			if !s.Supports(store.TypeDynDB) {
				continue
			}

			dyns[k] = s.DynDB()
		}

		b.app.dynHub = dynhub.New(b.app, dyns)
	}

	b.app.sockd = sockcore.New()
	b.app.coreHub = corehub.New(b.coreDB, b.app.sockd, b.cPlane)
	b.app.pacman = pacman.New(b.app, b.repos)

	return nil
}

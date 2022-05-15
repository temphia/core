package app

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/backend/server/btypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/temphia/temphia/backend/server/btypes/store"
)

type App struct {
	nodeId  string
	devmode bool

	registry       interface{}
	controlPlane   service.ControlPlane
	fencer         service.Fencer
	engine         rtypes.Engine
	sockd          service.SockCore
	coreHub        store.CoreHub
	plugKV         store.PlugStateKV
	signer         service.Signer
	cabinetHub     store.CabinetHub
	pacman         service.Pacman
	dynHub         store.DynHub
	data           store.AssetStore
	server         btypes.Server
	rootLogger     *zerolog.Logger
	rootController interface{}
}

func (a *App) Run() error {

	err := a.engine.Run()
	if err != nil {
		return err
	}

	return a.server.Listen()
}

func (a *App) NodeId() string {
	return a.nodeId
}

func (a *App) DevMode() bool { return a.devmode }

func (a *App) ControlPlane() interface{}   { return a.controlPlane }
func (a *App) Fencer() interface{}         { return a.fencer }
func (a *App) Engine() interface{}         { return a.engine }
func (a *App) Sockd() interface{}          { return a.sockd }
func (a *App) CoreHub() store.CoreHub      { return a.coreHub }
func (a *App) PlugKV() store.PlugStateKV   { return a.plugKV }
func (a *App) Signer() interface{}         { return a.signer }
func (a *App) Cabinet() store.CabinetHub   { return a.cabinetHub }
func (a *App) Pacman() interface{}         { return a.pacman }
func (a *App) DynHub() store.DynHub        { return a.dynHub }
func (a *App) Data() store.AssetStore      { return a.data }
func (a *App) Server() btypes.Server       { return a.server }
func (a *App) RootController() interface{} { return a.rootController }
func (a *App) RootLogger() *zerolog.Logger { return a.rootLogger }
func (a *App) Registry() interface{}       { return a.registry }

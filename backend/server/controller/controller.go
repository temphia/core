package controller

import (
	"github.com/temphia/temphia/backend/server/app/config"
	"github.com/temphia/temphia/backend/server/btypes"
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/temphia/temphia/backend/server/controller/admin"
	"github.com/temphia/temphia/backend/server/controller/auth"
	"github.com/temphia/temphia/backend/server/controller/basic"
	"github.com/temphia/temphia/backend/server/controller/cabinet"
	"github.com/temphia/temphia/backend/server/controller/dtable"
	"github.com/temphia/temphia/backend/server/controller/operator"
)

// Controller is Parent controller which holds all and inits all
// actual controller, where as controller mostly sits on top of
// api/routes, it gets all claim and params parsed from routes. It
// checks permisission bashed on claim and calls lower level services
// to finish the operations [ api/routes => controller => services => db/providers]

type RootController struct {
	cAdmin    *admin.Controller
	cAuth     *auth.Controller
	cBasic    *basic.Controller
	cCabinet  *cabinet.Controller
	cDtable   *dtable.Controller
	cOperator *operator.Controller
}

func New(_app btypes.App, config *config.AppConfig) *RootController {
	cplane := _app.ControlPlane().(service.ControlPlane)
	corehub := _app.CoreHub()
	pacman := _app.Pacman().(service.Pacman)
	signer := _app.Signer().(service.Signer)
	cab := _app.Cabinet()
	fencer := _app.Fencer().(service.Fencer)

	return &RootController{
		cAdmin:    admin.New(pacman, cplane, corehub, signer),
		cAuth:     auth.New(corehub, _app.Fencer().(service.Fencer), signer),
		cBasic:    basic.New(corehub, cab, _app.DynHub(), pacman),
		cCabinet:  cabinet.New(fencer, cab, signer),
		cDtable:   dtable.New(fencer, _app.DynHub(), cab, signer),
		cOperator: operator.New(corehub, signer, config),
	}
}

func (c *RootController) AdminController() *admin.Controller       { return c.cAdmin }
func (c *RootController) AuthController() *auth.Controller         { return c.cAuth }
func (c *RootController) BasicController() *basic.Controller       { return c.cBasic }
func (c *RootController) CabinetController() *cabinet.Controller   { return c.cCabinet }
func (c *RootController) DtableController() *dtable.Controller     { return c.cDtable }
func (c *RootController) OperatorController() *operator.Controller { return c.cOperator }

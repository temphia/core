package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/app/config"
	"github.com/temphia/core/backend/server/app/routes/site"
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/temphia/core/backend/server/controller"
	"github.com/temphia/core/backend/server/controller/admin"
	"github.com/temphia/core/backend/server/controller/auth"
	"github.com/temphia/core/backend/server/controller/basic"
	"github.com/temphia/core/backend/server/controller/cabinet"
	"github.com/temphia/core/backend/server/controller/dtable"
	"github.com/temphia/core/backend/server/controller/operator"
)

type R struct {
	app            btypes.App
	rootController *controller.RootController
	cAdmin         *admin.Controller
	cAuth          *auth.Controller
	cBasic         *basic.Controller
	cCabinet       *cabinet.Controller
	cDtable        *dtable.Controller
	cOperator      *operator.Controller
	siteValidator  func(c *gin.Context) error
	data           store.AssetStore
	assetFS        http.FileSystem
	engine         rtypes.Engine
	signer         service.Signer
	sockd          service.SockCore
	sitemanager    site.Manager
}

func New(_app btypes.App, config *config.AppConfig) *R {
	ctrl := controller.New(_app, config)
	return &R{
		app:            _app,
		rootController: ctrl,
		cAdmin:         ctrl.AdminController(),
		cAuth:          ctrl.AuthController(),
		cBasic:         ctrl.BasicController(),
		cCabinet:       ctrl.CabinetController(),
		cDtable:        ctrl.DtableController(),
		cOperator:      ctrl.OperatorController(),
		siteValidator:  nil,
		signer:         _app.Signer().(service.Signer),
		data:           _app.Data(),
		assetFS:        http.FS(_app.Data().AssetAdapter()),
		engine:         _app.Engine().(rtypes.Engine),
		sockd:          _app.Sockd().(service.SockCore),
		sitemanager:    site.NewManager(_app.Cabinet(), _app.CoreHub()),
	}

}

func (r *R) RootIndex(ctx *gin.Context) {
	r.sitemanager.ServeIndex(ctx)
}

func (r *R) NoRoute(c *gin.Context) {
	curPath := c.Request.URL.Path

	if strings.Contains(curPath, "/assets/") {
		pp.Println("@@no_path =>", curPath)

		paths := strings.Split(curPath, "/assets/")
		c.FileFromFS(paths[1], r.assetFS)
		return
	}

	if strings.HasPrefix(curPath, "/console") {
		// fixme => server console root file
		c.Redirect(http.StatusFound, "/console")
	}

	r.sitemanager.ServeAny(curPath, c)
}

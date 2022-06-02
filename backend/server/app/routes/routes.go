package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/core/backend/server/app/config"
	"github.com/temphia/core/backend/server/app/routes/hoster"
	"github.com/temphia/core/backend/server/services/sockdhub"

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
	sockdhub       sockdhub.SockdHub

	sitemanager hoster.Manager
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
		sockdhub:       *sockdhub.New(_app.Sockd().(service.SockCore)),
		sitemanager:    hoster.NewManager(_app.Cabinet(), _app.CoreHub()),
	}

}

func (r *R) RootIndex(ctx *gin.Context) {
	r.sitemanager.ServeIndex(ctx)
}

func (r *R) StartIndex(ctx *gin.Context) {
	hoster.DefaultIndex(ctx)
}

func (r *R) NoRoute(c *gin.Context) {
	r.sitemanager.ServeAny(c.Request.URL.Path, c)
}

func (r *R) ServeAssets(c *gin.Context) {
	c.FileFromFS(c.Param("file"), r.assetFS)
}

func (r *R) ServePublic(c *gin.Context) {

}

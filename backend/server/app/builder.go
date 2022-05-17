package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/temphia/core/backend/server/app/config"
	"github.com/temphia/core/backend/server/app/routes"
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/btypes/store"
	"github.com/temphia/core/backend/server/helpers/ginlogger"
	"github.com/temphia/core/backend/server/registry"

	"github.com/temphia/core/data"
)

type Builder struct {
	app        *App
	config     *config.AppConfig
	cPlane     service.ControlPlane
	middleware func(*gin.Context)
	data       store.AssetStore
	registery  *registry.Registry
	appLogger  *zerolog.Logger
	rtLogger   *zerolog.Logger

	ginEngine *gin.Engine

	beforeListen []func(btypes.App, *gin.Engine) error
	stores       map[string]store.Store
	repos        map[string]store.Repository
	coreDB       store.CoreDB
}

func WithApp(app *App) func(builder *Builder) {
	return func(builder *Builder) {
		builder.app = app
	}
}

// with app config
func WithConfig(conf *config.AppConfig) func(builder *Builder) {
	return func(builder *Builder) {
		builder.config = conf
	}
}

// with dev app config
func WithDevConfig() func(builder *Builder) {
	return func(builder *Builder) {
		builder.config = config.DefaultUnsafeDev()
	}
}

// with global gin middleware
func WithControlPlane(cp service.ControlPlane) func(builder *Builder) {
	return func(builder *Builder) {
		builder.cPlane = cp
	}
}

// with global gin middleware
func WithMiddleware(mw func(*gin.Context)) func(builder *Builder) {
	return func(builder *Builder) {
		builder.middleware = mw
	}
}

// with registry, if not provided global registry is cloned
func WithRegistry(registery *registry.Registry) func(builder *Builder) {
	return func(builder *Builder) {
		builder.registery = registery
	}
}

// application/default logger
func WithAppLogger(logger *zerolog.Logger) func(builder *Builder) {
	return func(builder *Builder) {
		builder.appLogger = logger
	}
}

// runtime logger
func WithRTLogger(logger *zerolog.Logger) func(builder *Builder) {
	return func(builder *Builder) {
		builder.rtLogger = logger
	}
}

func WithBeforeListen(hook func(btypes.App, *gin.Engine) error) func(builder *Builder) {
	return func(builder *Builder) {
		builder.beforeListen = append(builder.beforeListen, hook)
	}
}

func WithStore(name string, s store.Store) func(builder *Builder) {
	return func(builder *Builder) {
		builder.stores[name] = s
	}
}

func WithEngine(e *gin.Engine) func(builder *Builder) {
	return func(builder *Builder) {
		builder.ginEngine = e
	}
}

// NewBuilder creates builder with options
func NewBuilder(options ...func(*Builder)) *Builder {
	builder := &Builder{
		beforeListen: make([]func(btypes.App, *gin.Engine) error, 0),
		stores:       make(map[string]store.Store),
		repos:        make(map[string]store.Repository),
	}
	for _, opt := range options {
		opt(builder)
	}

	builder.applyDefault()

	return builder
}

func (b *Builder) applyDefault() {

	if b.cPlane == nil {
		panic("empty control plane")
	}

	if b.app == nil {
		b.app = &App{
			controlPlane: b.cPlane,
		}
	}

	if b.config == nil {
		panic("empty config")
	}

	//	b.app.nodeId = b.config.AppName
	// fixme => get node id from somewhere

	if b.data == nil {
		b.data = data.DefaultNew()
	}

	b.app.data = b.data

	if b.registery == nil {
		b.registery = registry.New(true)
	}

	if b.appLogger == nil {
		root := zerolog.New(zerolog.NewConsoleWriter())
		b.appLogger = &root
	}

	if b.rtLogger == nil {
		rtlogger := b.appLogger.With().Str("service", "runtime").Logger()
		b.rtLogger = &rtlogger
	}

	b.app.rootLogger = b.appLogger
}

func (b *Builder) Build() error {

	err := b.init()
	if err != nil {
		return err
	}

	err = b.buildComponents()
	if err != nil {
		return err
	}

	err = b.cPlane.Inject(b.app, b.config)
	if err != nil {
		return err
	}

	if b.ginEngine == nil {
		b.ginEngine = gin.New()
		gin.SetMode(gin.DebugMode)
		b.ginEngine.Use(ginlogger.Logger(b.appLogger, "GIN_APP"), gin.Recovery())
	}

	if b.middleware != nil {
		b.ginEngine.Use(b.middleware)
	}

	b.app.server = &Server{
		app:       b.app,
		routes:    routes.New(b.app, b.config),
		ginEngine: b.ginEngine,
		config:    b.config,
	}

	b.app.server.BindRoutes()

	err = b.app.controlPlane.Start()
	if err != nil {
		return err
	}

	for _, hook := range b.beforeListen {
		err := hook(b.app, b.ginEngine)
		if err != nil {
			return err
		}
	}

	b.app.cabinetHub.Start(b.cPlane.GetEventBus())
	return nil

}

func (b *Builder) init() error {
	errs := b.config.Check()
	if len(errs) > 0 {
		return errs[0]
	}

	return b.config.ApplyDefault("dev")
}

func (b *Builder) App() *App {
	return b.app
}

func (b *Builder) RunApp() error {
	return b.app.Run()
}

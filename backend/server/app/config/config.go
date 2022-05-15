package config

import (
	"crypto/rand"
	"errors"

	"github.com/rs/xid"
)

type AppConfig struct {
	AppName           string          `json:"app_name,omitempty"`
	ServerPort        string          `json:"http_port,omitempty"`
	MasterKeyType     string          `json:"master_key_type,omitempty"`
	MasterKey         string          `json:"master_key,omitempty"`
	Runtime           *RuntimeOptions `json:"runtime_options,omitempty"`
	StoreSources      []*StoreSource  `json:"db_sources,omitempty"`
	StoreOptions      *StoreOptions   `json:"store_options,omitempty"`
	BprintSources     []*StoreSource  `json:"bprint_sources,omitempty"`
	StartUpOptions    *StartUpOptions `json:"start_up_options,omitempty"`
	NodeOptions       *NodeOptions    `json:"node_options,omitempty"`
	OperatorName      string          `json:"op_name,omitempty"`
	OperatorPassword  string          `json:"op_password,omitempty"`
	SkipMasterOpToken string          `json:"skip_master_op_token,omitempty"`
}

type NodeOptions struct {
	Name          string            `json:"name,omitempty"`
	IdendityKey   string            `json:"idendity_key,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	DataOverrides map[string]string `json:"data_override,omitempty"`
}

type StoreOptions struct {
	DefaultCabinet string `json:"default_cabinet,omitempty"`
	DefaultCoreDB  string `json:"default_coredb,omitempty"`
}

type StartUpOptions struct {
	TenantName        string `json:"tenant_name,omitempty"`
	TenantSlug        string `json:"tenant_slug,omitempty"`
	SuperPassword     string `json:"super_password,omitempty"`
	NotPrintPassword  string `json:"not_print_pass,omitempty"`
	LocalDumpPassword bool   `json:"local_dump_pass,omitempty"`
	SkipIfExists      bool   `json:"skip_if_exists,omitempty"`
	AutoExitOnInit    bool   `json:"auto_exit_on_init,omitempty"`
	AutoDumpConfig    bool   `json:"auto_dump_config,omitempty"`
	AutoMigrate       bool   `json:"auto_migrate,omitempty"`
}

type StoreSource struct {
	Name     string                 `json:"name,omitempty"`
	Vendor   string                 `json:"vendor,omitempty"`
	Provider string                 `json:"provider,omitempty"`
	HostPath string                 `json:"host_path,omitempty"`
	User     string                 `json:"user,omitempty"`
	Password string                 `json:"password,omitempty"`
	Target   string                 `json:"target,omitempty"`
	Port     string                 `json:"port,omitempty"`
	Features []string               `json:"features,omitempty"`
	Options  map[string]interface{} `json:"options,omitempty"`
}

type RuntimeOptions struct {
	MaxWorker       int                    `json:"max_workers,omitempty"`
	MaxWorkerTenant int                    `json:"max_worker_tenant,omitempty"`
	TmpFolder       string                 `json:"tmp_folder,omitempty"`
	ExecutorOptions map[string]interface{} `json:"executors,omitempty"`
	ModulesOptions  map[string]interface{} `json:"modules,omitempty"`
	EnableMagicDial bool                   `json:"magic_dial,omitempty"`
}

func (c *AppConfig) Check() []error {
	errs := make([]error, 0)

	if c.AppName == "" {
		errs = append(errs, errors.New("Empty App Name"))
	}

	if c.MasterKeyType == "" {
		errs = append(errs, errors.New("Empty Master Keytype"))
	}

	if c.StoreOptions.DefaultCabinet == "" {
		errs = append(errs, errors.New("Empty Default Cabinet"))
	}

	if c.StoreOptions.DefaultCoreDB == "" {
		errs = append(errs, errors.New("Empty Default Coredb"))
	}

	return errs
}

func (c *AppConfig) ApplyDefault(mode string) error {

	if c.NodeOptions == nil {
		c.NodeOptions.Name = xid.New().String()
		c.NodeOptions.Tags = []string{}
		c.NodeOptions.DataOverrides = make(map[string]string)
	}

	if c.BprintSources == nil {
		c.BprintSources = make([]*StoreSource, 0, 3)
	}

	c.BprintSources = append(c.BprintSources, &StoreSource{
		Name:     "Official",
		Provider: "gitlab",
		HostPath: "https://gitlab.com/temphia/blueprint_store/",
		Options:  nil,
	})

	c.BprintSources = append(c.BprintSources, &StoreSource{
		Name:     "embed",
		Provider: "embed",
		Options:  nil,
	})

	c.BprintSources = append(c.BprintSources, &StoreSource{
		Name:     "local",
		Provider: "local",
		HostPath: "./data/devrepo/",
		Options:  nil,
	})

	c.BprintSources = append(c.BprintSources, &StoreSource{
		Name:     "Community",
		Provider: "gitlab",
		HostPath: "https://gitlab.com/temphia/blueprint_store/",
		Options:  nil,
	})

	return nil
}

func DefaultUnsafeDev() *AppConfig {
	randKey := make([]byte, 0, 10)
	rand.Read(randKey)

	return &AppConfig{
		AppName:       "Dev App",
		ServerPort:    ":4000",
		MasterKeyType: "simple_key",
		MasterKey:     string(randKey),
		Runtime: &RuntimeOptions{
			MaxWorker:       10,
			MaxWorkerTenant: 10,
			TmpFolder:       "",
			ExecutorOptions: map[string]interface{}{},
			ModulesOptions:  map[string]interface{}{},
			EnableMagicDial: false,
		},
		StoreSources: []*StoreSource{
			{
				Name:     "default_db",
				Provider: "postgres",
				Vendor:   "postgres",
				HostPath: "localhost",
				Target:   "temphia",
				User:     "temphia",
				Password: "temphia123",
				Port:     "7032",
				Features: []string{"core_db", "state_db", "dyn_db"},
				Options:  map[string]interface{}{},
			},
			{
				Name:     "default_fs",
				Provider: "local_fs",
				HostPath: "./tmp/files",
			},
		},
		StoreOptions: &StoreOptions{
			DefaultCabinet: "default_fs",
			DefaultCoreDB:  "default_db",
		},
		BprintSources:  []*StoreSource{},
		StartUpOptions: &StartUpOptions{},
		NodeOptions: &NodeOptions{
			Name:          "Node 1",
			IdendityKey:   "",
			Tags:          []string{"worker"},
			DataOverrides: map[string]string{},
		},
		OperatorName:      "operator",
		OperatorPassword:  "oops123",
		SkipMasterOpToken: "",
	}
}

package entities

type Plug struct {
	Id        string     `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Executor  string     `json:"executor,omitempty" db:"executor,omitempty"`
	Live      bool       `json:"live,omitempty" db:"live,omitempty"`
	Dev       bool       `json:"dev,omitempty"  db:"dev,omitempty"`
	Owner     string     `json:"owner,omitempty"  db:"owner,omitempty"`
	BprintId  string     `json:"bprint_id,omitempty"  db:"bprint_id,omitempty"`
	Handlers  JsonStrMap `json:"handlers,omitempty" db:"handlers,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type Agent struct {
	Id           string     `json:"id,omitempty" db:"id,omitempty"`
	Name         string     `json:"name,omitempty" db:"name,omitempty"`
	Type         string     `json:"type,omitempty" db:"type,omitempty"`
	InvokePolicy string     `json:"invoke_policy,omitempty" db:"invoke_policy,omitempty"`
	PlugID       string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	Resources    JsonStrMap `json:"resources,omitempty" db:"resources,omitempty"`
	ServeFiles   JsonStrMap `json:"serve_files,omitempty" db:"serve_files,omitempty"`

	EntryName   string `json:"entry_name,omitempty" db:"entry_name,omitempty"`
	EntryScript string `json:"entry_script,omitempty" db:"entry_script,omitempty"`
	EntryStyle  string `json:"entry_style,omitempty" db:"entry_style,omitempty"`
	ExecLoader  string `json:"exec_loader,omitempty" db:"exec_loader,omitempty"`

	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type PlugLink struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	FromPlug  string     `json:"from_plug_id,omitempty" db:"from_plug_id,omitempty"`
	FromAgent string     `json:"from_agent_id,omitempty" db:"from_agent_id,omitempty"`
	ToPlug    string     `json:"to_plug_id,omitempty" db:"to_plug,omitempty"`
	ToAgent   string     `json:"to_agent_id,omitempty" db:"to_agent,omitempty"`
	ToHandler string     `json:"to_handler,omitempty" db:"to_handler,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type PlugExtension struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Plug      string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	Agent     string     `json:"agent_id,omitempty" db:"agent_id,omitempty"`
	RefFile   string     `json:"ref_file,omitempty" db:"ref_file,omitempty"`
	BprintId  string     `json:"bprint_id,omitempty" db:"bprint_id,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

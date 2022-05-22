package entities

type Tenant struct {
	Name         string     `json:"name,omitempty" db:"name,omitempty"`
	Slug         string     `json:"slug,omitempty" db:"slug,omitempty"`
	OrgBio       string     `json:"org_bio,omitempty" db:"org_bio,omitempty"`
	RootPlug     string     `json:"root_plug,omitempty" db:"root_plug,omitempty"`
	RootAgent    string     `json:"root_agent,omitempty" db:"root_agent,omitempty"`
	SmtpUser     string     `json:"smtp_user,omitempty" db:"smtp_user,omitempty"`
	SmtpPass     string     `json:"smtp_pass,omitempty" db:"smtp_pass,omitempty"`
	MasterSecret string     `json:"master_secret,omitempty" db:"master_secret,omitempty"`
	DisableP2P   bool       `json:"disable_p2p,omitempty" db:"disable_p2p,omitempty"`
	ExtraMeta    JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type TenantDomain struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	About     string     `json:"about,omitempty" db:"about,omitempty"`
	CabSource string     `json:"cab_source,omitempty" db:"cab_source,omitempty"`
	Folder    string     `json:"folder,omitempty" db:"folder,omitempty"`
	TspPlug   string     `json:"tsp_plug,omitempty" db:"tsp_plug,omitempty"`
	TspAgent  string     `json:"tsp_agent,omitempty" db:"tsp_agent,omitempty"`
	SMTPUser  string     `json:"smtp_user,omitempty" db:"smtp_user,omitempty"`
	SMTPPass  string     `json:"smtp_pass,omitempty" db:"smtp_pass,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type DomainWidget struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Plug      string     `json:"plug,omitempty" db:"plug,omitempty"`
	Agent     string     `json:"agent,omitempty" db:"agent,omitempty"`
	ExecMeta  JsonMap    `json:"exec_meta,omitempty" db:"exec_meta,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

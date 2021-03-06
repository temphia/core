package entities

type Tenant struct {
	Name         string     `json:"name,omitempty" db:"name,omitempty"`
	Slug         string     `json:"slug,omitempty" db:"slug,omitempty"`
	OrgBio       string     `json:"org_bio,omitempty" db:"org_bio,omitempty"`
	RootPlug     string     `json:"root_plug,omitempty" db:"root_plug,omitempty"`
	RootAgent    string     `json:"root_agent,omitempty" db:"root_agent,omitempty"`
	SmtpUser     string     `json:"smtp_user,omitempty" db:"smtp_user,omitempty"`
	SmtpPassword string     `json:"smtp_password,omitempty" db:"smtp_password,omitempty"`
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
	Widgets   JsonStrMap `json:"widgets,omitempty" db:"widgets,omitempty"` // fixme => json_int_map
	SMTPUser  string     `json:"smtp_user,omitempty" db:"smtp_user,omitempty"`
	SMTPPass  string     `json:"smtp_pass,omitempty" db:"smtp_pass,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type TenantWizard struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Plug      string     `json:"plug,omitempty" db:"plug,omitempty"`
	Agent     string     `json:"agent,omitempty" db:"agent,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

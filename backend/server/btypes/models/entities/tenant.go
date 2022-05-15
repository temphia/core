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
	Id           int64      `json:"id,omitempty" db:"id,omitempty"`
	Name         string     `json:"name,omitempty" db:"name,omitempty"`
	About        string     `json:"about,omitempty" db:"about,omitempty"`
	Folder       string     `json:"folder,omitempty" db:"folder,omitempty"`
	IndexFile    string     `json:"index_file,omitempty" db:"index_file,omitempty"`
	RenderType   string     `json:"render_type,omitempty" db:"render_type,omitempty"`
	RenderOption JsonStrMap `json:"render_options,omitempty" db:"render_options,omitempty"`
	SMTPUser     string     `json:"smtp_user,omitempty" db:"smtp_user,omitempty"`
	SMTPPass     string     `json:"smtp_pass,omitempty" db:"smtp_pass,omitempty"`
	ExtraMeta    JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

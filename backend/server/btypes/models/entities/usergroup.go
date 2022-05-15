package entities

type UserGroup struct {
	Name      string     `json:"name,omitempty" db:"name"`
	Slug      string     `json:"slug,omitempty" db:"slug"`
	Icon      string     `json:"icon,omitempty" db:"icon"`
	TenantID  string     `json:"tenant_id,omitempty" db:"tenant_id"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type UserGroupUpdate struct {
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Icon      string     `json:"icon,omitempty" db:"icon,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

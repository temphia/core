package entities

type BPrint struct {
	ID           string    `json:"id,omitempty" db:"id,omitempty"`
	Name         string    `json:"name,omitempty" db:"name,omitempty"`
	Slug         string    `json:"slug,omitempty" db:"name,omitempty"`
	Type         string    `json:"type,omitempty" db:"type,omitempty"`
	SubType      string    `json:"sub_type,omitempty" db:"sub_type,omitempty"`
	InlineSchema string    `json:"inline_schema,omitempty"  db:"inline_schema,omitempty"`
	Description  string    `json:"description,omitempty" db:"description,omitempty"`
	Icon         string    `json:"icon,omitempty" db:"icon,omitempty"`
	Source       string    `json:"source,omitempty" db:"source,omitempty"`
	TenantID     string    `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	Tags         JsonArray `json:"tags,omitempty" db:"tags,omitempty"`
	Files        JsonArray `json:"files,omitempty" db:"files,omitempty"`
	ExtraMeta    JsonMap   `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

func (b *BPrint) Combine(data map[string]interface{}) {
	setString("description", &b.Description, data)
	setString("type", &b.Type, data)
	setString("inline_schema", &b.InlineSchema, data)
	setString("slug", &b.Slug, data)
	setString("sub_type", &b.SubType, data)
}

func setString(name string, pstr *string, data map[string]interface{}) {
	s, ok := data[name]
	if !ok {
		return
	}
	str, ok := s.(string)
	if !ok {
		return
	}
	pstr = &str
}

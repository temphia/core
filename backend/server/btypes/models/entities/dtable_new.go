package entities

type NewTableGroup struct {
	Name        string      `json:"name,omitempty" yaml:"name,omitempty"`
	Slug        string      `json:"slug,omitempty" yaml:"slug,omitempty"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Tables      []*NewTable `json:"tables,omitempty" yaml:"tables,omitempty"`
	ExecOrder   []string    `json:"exec_order,omitempty" yaml:"exec_order,omitempty"`
}

type NewTable struct {
	Name          string         `json:"name,omitempty" yaml:"name,omitempty"`
	Slug          string         `json:"slug,omitempty" yaml:"slug,omitempty"`
	Description   string         `json:"description,omitempty" yaml:"description,omitempty"`
	Icon          string         `json:"icon,omitempty" yaml:"icon,omitempty"`
	MainColumn    string         `json:"main_column,omitempty" yaml:"main_column,omitempty"`
	Columns       []*NewColumn   `json:"columns,omitempty" yaml:"columns,omitempty"`
	Indexes       []Index        `json:"indexes,omitempty" yaml:"indexes,omitempty"`
	UniqueIndexes []Index        `json:"unique_indexes,omitempty" yaml:"unique_indexes,omitempty"`
	FTSIndex      *FTSIndex      `json:"fts_index,omitempty" yaml:"fts_index,omitempty"`
	ColumnRef     []*ColumnFKRef `json:"column_refs,omitempty" yaml:"column_refs,omitempty"`
	DeletedAt     bool           `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
	Views         []View         `json:"views,omitempty" yaml:"views,omitempty"`
	SeedData      SeedData       `json:"seed_data,omitempty" yaml:"seed_data,omitempty"`
}

type NewColumn struct {
	Name        string `json:"name,omitempty" db:"name,omitempty" yaml:"name,omitempty"`
	Slug        string `json:"slug,omitempty" db:"slug,omitempty" yaml:"slug,omitempty"`
	Ctype       string `json:"ctype,omitempty" db:"ctype,omitempty" yaml:"ctype,omitempty"`
	Description string `json:"description,omitempty" db:"description,omitempty" yaml:"description,omitempty"`
	Icon        string `json:"icon,omitempty" db:"icon,omitempty" yaml:"icon,omitempty"`

	Options       []string `json:"options,omitempty" db:"options,omitempty" yaml:"options,omitempty"`
	NotNullable   bool     `json:"not_nullable,omitempty" db:"not_nullable,omitempty" yaml:"not_nullable,omitempty"`
	Pattern       string   `json:"pattern,omitempty" db:"pattern,omitempty" yaml:"pattern,omitempty"`
	StrictPattern bool     `json:"strict_pattern,omitempty" db:"strict_pattern,omitempty" yaml:"strict_pattern,omitempty"`
}

type Index struct {
	Mtype string   `json:"mtype,omitempty" yaml:"mtype,omitempty"`
	Slug  string   `json:"slug,omitempty" yaml:"slug,omitempty"`
	Spans []string `json:"spans" yaml:"spans"`
}

type FTSIndex struct {
	Type        string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Slug        string                 `json:"slug,omitempty" yaml:"slug,omitempty"`
	ColumnSpans []string               `json:"spans" yaml:"spans"`
	Options     map[string]interface{} `json:"options" yaml:"options"`
}

type ColumnFKRef struct {
	Slug     string   `json:"slug,omitempty" yaml:"slug,omitempty"`
	Type     string   `json:"type,omitempty" yaml:"type,omitempty"`
	Target   string   `json:"target,omitempty" yaml:"target,omitempty"`
	FromCols []string `json:"from_cols,omitempty" yaml:"from_cols,omitempty"`
	ToCols   []string `json:"to_cols,omitempty" yaml:"to_cols,omitempty"`
	RefCopy  string   `json:"ref_copy,omitempty" yaml:"ref_copy,omitempty"`
}

type SeedData struct {
	Data         []map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`
	LinkedImages []string                 `json:"linked_images,omitempty" yaml:"linked_images,omitempty"`
}

func (m *NewTableGroup) To(tenantId string) *TableGroup {
	return &TableGroup{
		Name:        m.Name,
		Slug:        m.Slug,
		Description: m.Description,
		SourceDb:    "default",
		TenantID:    tenantId,
	}
}

func (m *NewTable) To(tenantId, gslug string) *Table {
	return &Table{
		Name:        m.Name,
		Slug:        m.Slug,
		Description: m.Description,
		Icon:        m.Icon,
		GroupID:     gslug,
		TenantID:    tenantId,
		MainColumn:  "",
	}
}

func (m *NewColumn) To(tenantId, gslug, tslug string) *Column {
	return &Column{
		Name:          m.Name,
		Slug:          m.Slug,
		Ctype:         m.Ctype,
		Description:   m.Description,
		GroupID:       gslug,
		Icon:          m.Icon,
		Options:       m.Options,
		OrderID:       0,
		Pattern:       m.Pattern,
		StrictPattern: m.StrictPattern,
		TableID:       tslug,
		TenantID:      tenantId,
		RefId:         "",
		RefType:       "",
		RefTarget:     "",
		RefObject:     "",
		RefCopy:       "",
		ExtraMeta:     nil,
	}
}

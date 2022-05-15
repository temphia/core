package entities

type Permission struct {
	ID            int64  `json:"id,omitempty" db:"id,omitempty"`
	ObjectType    string `json:"object_type,omitempty" db:"object_type,omitempty"`
	ObjectSubType string `json:"object_sub_type,omitempty" db:"object_sub_type,omitempty"`
	ObjectPath    string `json:"object_path,omitempty" db:"object_path,omitempty"`
	ObjectAction  string `json:"object_action,omitempty" db:"object_action,omitempty"`
	ObjectData    string `json:"object_data,omitempty" db:"object_data,omitempty"`
	RoleID        string `json:"role_id,omitempty" db:"role_id,omitempty"`
	TenantID      string `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type Role struct {
	Id       string `json:"id,omitempty" db:"id,omitempty"`
	TenantID string `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type UserRole struct {
	RoleId   string `json:"role_id,omitempty" db:"role_id,omitempty"`
	TenantID string `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	Username string `json:"username,omitempty" db:"username,omitempty"`
}

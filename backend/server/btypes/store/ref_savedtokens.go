package store

import "time"

type SavedToken struct {
	Id        string     `json:"name,omitempty" db:"name,omitempty"`
	Type      string     `json:"type,omitempty" db:"type,omitempty"`
	UserId    string     `json:"user_id,omitempty" db:"user_id,omitempty"`
	UserGroup string     `json:"user_group,omitempty" db:"user_group,omitempty"`
	Target    string     `json:"target,omitempty" db:"target,omitempty"`
	Payload   string     `json:"payload,omitempty" db:"payload,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExpiresOn *time.Time `json:"expires_on,omitempty" db:"expires_on,omitempty"`
}

type SavedTokenQuery struct {
	UserId    string `json:"user_id,omitempty"`
	UserGroup string `json:"user_group,omitempty"`
	Target    string `json:"target,omitempty"`
}

type SavedTokens interface {
	NewSavedToken(data *SavedToken) error
	QuerySavedToken(query *SavedTokenQuery) ([]*SavedToken, error)
	GetSavedToken(tenatId string, id string) (*SavedToken, error)
	DeleteSavedToken(tenatId string, id string) error
}

package claim

import (
	"time"

	"github.com/rs/xid"
)

type User struct {
	TenentID   string            `json:"tenent_id,omitempty"`
	UserID     string            `json:"user_id,omitempty"`
	UserGroup  string            `json:"user_group,omitempty"`
	Type       string            `json:"type,omitempty"`
	Expiry     int64             `json:"expiry,omitempty"`
	DeviceID   string            `json:"device_id,omitempty"`
	Scopes     []Scope           `json:"scopes,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

func NewUserDevice(tenantId, userId, groupId string) *User {
	return &User{
		TenentID:   tenantId,
		Expiry:     time.Now().Unix(),
		Type:       "user_device",
		UserID:     userId,
		UserGroup:  groupId,
		DeviceID:   xid.New().String(),
		Scopes:     []Scope{},
		Attributes: make(map[string]string),
	}
}

func NewUserLogged(tenantId, userId, groupId string) *User {
	return &User{
		TenentID:   tenantId,
		Expiry:     time.Now().Unix(),
		Type:       "user_logged",
		UserID:     userId,
		UserGroup:  groupId,
		DeviceID:   xid.New().String(),
		Scopes:     []Scope{},
		Attributes: make(map[string]string),
	}
}

func (u *User) Valid() error { return nil }

func (u *User) IsSuperAdmin() bool {
	return u.UserGroup == "super_admin"
}

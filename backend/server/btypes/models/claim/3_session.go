package claim

import (
	"time"

	"github.com/rs/xid"
)

type RBACMode string

const (
	RBACSkip RBACMode = "rbac_skip"
	RBACJit  RBACMode = "rbac_jit"
	RBACAot  RBACMode = "rbac_aot"
)

type Session struct {
	TenentID    string            `json:"-"`
	UserID      string            `json:"user,omitempty"`
	UserGroup   string            `json:"group,omitempty"`
	Type        string            `json:"type,omitempty"`
	Expiry      int64             `json:"expiry,omitempty"`
	SessionID   string            `json:"session_id,omitempty"`
	DeviceID    string            `json:"device_id,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	ServicePath []string          `json:"service_path,omitempty"`
	RBACMode    RBACMode          `json:"rbac_mode,omitempty"`
	Objects     []Object          `json:"objects,omitempty"`
}

type Object struct {
	Path          []string `json:"path,omitempty"`
	Actions       []string `json:"actions,omitempty"`
	WhitelistMode bool     `json:"whitelist_mode,omitempty"`
}

func NewSession(user, group, device string) *Session {
	return &Session{
		TenentID:   "",
		UserID:     user,
		UserGroup:  group,
		Type:       "session",
		Expiry:     time.Now().Unix(),
		SessionID:  xid.New().String(),
		DeviceID:   device,
		Attributes: make(map[string]string), Objects: []Object{},
	}
}

func (p *Session) Valid() error { return nil }

func (p *Session) SetAttr(key, value string) {
	p.Attributes[key] = value
}

func (u *Session) IsSuperAdmin() bool {
	return u.UserGroup == "super_admin"
}

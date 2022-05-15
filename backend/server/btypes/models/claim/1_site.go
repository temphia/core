package claim

import (
	"time"

	"github.com/rs/xid"
)

type Site struct {
	TenentID   string            `json:"tenent_id,omitempty"`
	ClaimType  string            `json:"type,omitempty"`
	Expiry     int64             `json:"expiry,omitempty"`
	XID        string            `json:"xid,omitempty"`
	SiteType   string            `json:"site_type,omitempty"`
	Origin     string            `json:"origin,omitempty"`
	Scopes     []*Scope          `json:"scopes,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

func NewSiteClaim(tenantId, origin string) *Site {
	return &Site{
		TenentID:   tenantId,
		Expiry:     time.Now().Unix(),
		ClaimType:  "site_claim",
		XID:        xid.New().String(),
		SiteType:   "admin_console",
		Origin:     origin,
		Scopes:     []*Scope{},
		Attributes: make(map[string]string),
	}
}

func (p *Site) Valid() error { return nil }

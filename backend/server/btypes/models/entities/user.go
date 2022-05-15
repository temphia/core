package entities

import (
	"database/sql/driver"
	"time"
)

type User struct {
	UserId    string    `json:"user_id,omitempty" db:"user_id"`
	FullName  string    `json:"full_name,omitempty" db:"full_name"`
	Email     string    `json:"email,omitempty" db:"email"`
	Bio       string    `json:"bio,omitempty" db:"bio"`
	GroupID   string    `json:"group_id,omitempty" db:"group_id"`
	Password  string    `json:"password,omitempty" db:"password"`
	TenantID  string    `json:"tenant_id,omitempty" db:"tenant_id"`
	PublicKey string    `json:"pub_key,omitempty" db:"pub_key"`
	Data      *UserData `json:"data,omitempty" db:"data"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	Active    bool      `json:"active,omitempty" db:"active"`
}

type UserData struct {
	// MFAEnabled         bool                   `json:"mfa_enabled,omitempty"`
	// MFAType            string                 `json:"mfa_type,omitempty"`
	// MFASeed            string                 `json:"mfa_seed,omitempty"`
	// EmailCode          string                 `json:"email_code,omitempty"`
	// AuthTokens         map[string]string      `json:"auth_token,omitempty"`
	// APNTokens          []string               `json:"apn_tokens,omitempty"`
	PendingPassChange  bool                   `json:"pending_pass_change,omitempty"`
	PendingEmailVerify bool                   `json:"pending_email_verify,omitempty"`
	Meta               map[string]interface{} `json:"meta,omitempty"`
}

func (ud *UserData) Value() (driver.Value, error) {
	return JSONDriverValue(ud)
}

func (ud *UserData) Scan(value interface{}) error {
	return JSONDriverScan(ud, value)
}

type UserUpdate struct {
	FullName  string   `json:"full_name,omitempty" db:"full_name,omitempty"`
	Email     string   `json:"email,omitempty" db:"email,omitempty"`
	GroupID   string   `json:"group_id,omitempty" db:"group_id,omitempty"`
	Password  string   `json:"password,omitempty" db:"password,omitempty"`
	PublicKey string   `json:"pub_key,omitempty" db:"pub_key,omitempty"`
	AuthType  string   `json:"auth_type,omitempty" db:"auth_type,omitempty"`
	Data      UserData `json:"data,omitempty" db:"data,omitempty"`
}

type UserInfo struct {
	UserId    string `json:"user_id,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	Bio       string `json:"bio,omitempty" db:"bio"`
	PublicKey string `json:"pub_key,omitempty"`
	Email     string `json:"email,omitempty"`
	Group     string `json:"group,omitempty"`
}

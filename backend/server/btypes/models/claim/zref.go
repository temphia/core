package claim

type UserPrekey struct {
	TenentID  string `json:"tenent_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	UserGroup string `json:"user_group,omitempty"`
	Type      string `json:"type,omitempty"`
}

package vmodels

type PlugData struct {
	Name       string `json:"name,omitempty"`
	PlugId     string `json:"plug_id,omitempty"`
	AgentId    string `json:"agent_id,omitempty"`
	ExecTicket string `json:"exec_ticket,omitempty"`
}

type UserData struct {
	Name  string `json:"name,omitempty"`
	Id    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Group string `json:"group,omitempty"`
}

type LoginRequest struct {
	TenantId     string `json:"tenant_id,omitempty"`
	UserIdendity string `json:"user_idendity,omitempty"`
	Password     string `json:"password,omitempty"`
	SiteToken    string `json:"site_token,omitempty"`
}

type LoginResponse struct {
	Token    string     `json:"token,omitempty"`
	Message  string     `json:"message,omitempty"`
	Console  string     `json:"console,omitempty"`
	Plugs    []PlugData `json:"plugs,omitempty"`
	UserData UserData   `json:"user_data,omitempty"`
}

type RefreshReq struct {
	Path      []string               `json:"path,omitempty"`
	UserToken string                 `json:"user_token,omitempty"`
	Options   map[string]interface{} `json:"options,omitempty"`
	OldToken  string                 `json:"old_token,omitempty"`
}

type RefreshResp struct {
	Token    string `json:"token,omitempty"`
	Expiry   string `json:"expiry,omitempty"`
	Message  string `json:"message,omitempty"`
	StatusOk bool   `json:"status_ok,omitempty"`
}

// type LoginPageOptions struct {
// 	PinnedGroup  string `json:"pgroup,omitempty"`
// 	PinnedTenant string `json:"ptenant,omitempty"`
// 	LoginMode    string `json:"loginmode,omitempty"`
// 	Provider     string `json:"provider,omitempty"`
// 	ProviderIcon string `json:"picon,omitempty"`
// }

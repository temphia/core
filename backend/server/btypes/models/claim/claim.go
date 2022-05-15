package claim

const (
	CTypeSite       = "site"
	CTypeUserLogged = "user_logged"
	CTypeUserDevice = "user_device"
	CTypeSession    = "session"
)

type Scope struct {
	Paths   []string `json:"paths,omitempty"`
	Actions []string `json:"actions,omitempty"`
}

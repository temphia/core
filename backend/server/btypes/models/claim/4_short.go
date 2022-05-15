package claim

type TicketCabinet struct {
	Expiry   int64  `json:"z"`
	Mode     string `json:"a,omitempty"` // mode is ?
	Folder   string `json:"b,omitempty"`
	Source   string `json:"c,omitempty"`
	Prefix   string `json:"d,omitempty"`
	DeviceId string `json:"e,omitempty"`
}

func (tc *TicketCabinet) Valid() error { return nil } // fixme => validate

/*
	you could pin list of file exactly or give prefix

	PinnedFiles  []string `json:"f,omitempty"`
	ListPrefix   string   `json:"g,omitempty"`
	GetPrefix    string   `json:"h,omitempty"`
	UploadPrefix string   `json:"i,omitempty"`

	AllowUpload   bool `json:"j,omitempty"`
	AllowList     bool `json:"k,omitempty"`
	AllowGet      bool `json:"l,omitempty"`
	AllowShare    bool `json:"m,omitempty"`
	MaxUploadSize int  `json:"n,omitempty"` // in KB // fixme => implement this


*/

// used by device pair

type ServerToken struct {
	PrimaryURL string   `json:"primary_url,omitempty"`
	ExtraURL   []string `json:"extra_url,omitempty"`
	TenentId   string   `json:"tenant_id,omitempty"`
	Token      string   `json:"token,omitempty"`
}

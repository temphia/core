package claim

type OperatorClaim struct {
	XID          string `json:"xid,omitempty"`
	ClaimType    string `json:"type,omitempty"`
	Expiry       int64  `json:"expiry,omitempty"`
	Origin       string `json:"origin,omitempty"`
	BindDeviceId string `json:"bind_device,omitempty"`
}

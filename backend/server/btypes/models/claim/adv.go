package claim

type Advisery struct {
	XID    string                 `json:"xid"`
	Type   string                 `json:"type"`
	Expiry int64                  `json:"expiry"`
	Data   map[string]interface{} `json:"data"`
}

func (p *Advisery) Valid() error { return nil }

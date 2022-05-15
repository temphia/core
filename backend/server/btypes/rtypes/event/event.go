package event

type Request struct {
	Id   string                 `json:"id,omitempty"`
	Type string                 `json:"type,omitempty"`
	Name string                 `json:"name,omitempty"`
	Vars map[string]interface{} `json:"vars,omitempty"`
	Data []byte                 `json:"data,omitempty"`
}

type Response struct {
	Vars    map[string]interface{} `json:"vars,omitempty"`
	Payload interface{}            `json:"payload,omitempty"`
}

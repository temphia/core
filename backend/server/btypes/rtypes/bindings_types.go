package rtypes

type Value struct {
	Value    string `json:"value,omitempty"`
	Audience string `json:"audience,omitempty"`
	Version  int64  `json:"version,omitempty"`
	TTL      int64  `json:"ttl,omitempty"`
}

type Resource struct {
	Name    string            `json:"name,omitempty"`
	Type    string            `json:"type,omitempty"`
	Payload string            `json:"payload,omitempty"`
	Meta    map[string]string `json:"meta,omitempty"`
}

type Connection struct {
	Name      string            `json:"name,omitempty"`
	Type      string            `json:"type,omitempty"`
	FromPlug  string            `json:"from_plug,omitempty"`
	FromAgent string            `json:"from_agent,omitempty"`
	ToPlug    string            `json:"to_plug,omitempty"`
	ToAgent   string            `json:"to_agent,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
}

type CabTicket struct {
	Prefix      string   `json:"prefix,omitempty"`
	PinnedFiles []string `json:"pinned_files,omitempty"`
	Operations  []string `json:"ops,omitempty"`
}

type HTTPRequest struct {
	Method  string            `json:"method,omitempty"`
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    []byte            `json:"body,omitempty"`
}

type HTTPResponse struct {
	SatusCode int                 `json:"status_code,omitempty"`
	Headers   map[string][]string `json:"headers,omitempty"`
	Json      bool                `json:"json,omitempty"`
	Body      []byte              `json:"body,omitempty"`
}

type SlotOption struct {
	Path  string                 `json:"path,omitempty"`
	Data  []byte                 `json:"data,omitempty"`
	Async bool                   `json:"async,omitempty"`
	Meta  map[string]interface{} `json:"meta,omitempty"`
}

type RenderOption struct {
	AsResponse bool                   `json:"as_response,omitempty"`
	Includes   []string               `json:"includes,omitempty"`
	Data       map[string]interface{} `json:"meta,omitempty"`
}

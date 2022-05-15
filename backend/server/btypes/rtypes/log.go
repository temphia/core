package rtypes

type LogQueryReq struct {
	PlugIds   []string `json:"plug_ids,omitempty"`
	AgentIds  []string `json:"agent_ids,omitempty"`
	RequetIds []string `json:"request_ids,omitempty"`
	LogIds    []string `json:"log_ids,omitempty"`
	DeviceIds []string `json:"device_ids,omitempty"`
	FromDate  int64    `json:"from_date,omitempty"`
}

type RuntimeLogger interface {
	Log(tenantId, plugId, agentId, reqid, logId string, message string)
	Query(tenantId string, req LogQueryReq) ([]string, error)
}

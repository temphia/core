package runtime

func (r *runtime) Log(tenantId, plugId, agentId, reqid, logId string, message string) {
	r.logger.Log().
		Str("tenant_id", tenantId).
		Str("plug_id", plugId).
		Str("agent_id", agentId).
		Str("req_id", reqid).
		Str("log_id", logId).
		Msg(message)
}

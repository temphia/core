package corehub

import "github.com/temphia/temphia/backend/server/btypes/models/entities"

// bprint

func (c *CoreHub) BprintNew(tenantId string, et *entities.BPrint) error {
	return c.coredb.BprintNew(tenantId, et)
}

func (c *CoreHub) BprintUpdate(tenantId, id string, data map[string]interface{}) error {
	return c.coredb.BprintUpdate(tenantId, id, data)
}

func (c *CoreHub) BprintGet(tenantId, id string) (*entities.BPrint, error) {
	return c.coredb.BprintGet(tenantId, id)
}

func (c *CoreHub) BprintDel(tenantId, id string) error {
	return c.coredb.BprintDel(tenantId, id)
}

func (c *CoreHub) BprintList(tenantId, group string) ([]*entities.BPrint, error) {
	return c.coredb.BprintList(tenantId, group)
}

// plug

func (c *CoreHub) PlugNew(tenantId string, pg *entities.Plug) error {
	return c.coredb.PlugNew(tenantId, pg)
}

func (c *CoreHub) PlugUpdate(tenantId string, id string, data map[string]interface{}) error {
	return c.coredb.PlugUpdate(tenantId, id, data)
}

func (c *CoreHub) PlugGet(tenantId, pid string) (*entities.Plug, error) {
	return c.coredb.PlugGet(tenantId, pid)
}

func (c *CoreHub) PlugDel(tenantId, pid string) error {
	return c.coredb.PlugDel(tenantId, pid)
}

func (c *CoreHub) PlugList(tenantId string) ([]*entities.Plug, error) {
	return c.coredb.PlugList(tenantId)
}

func (c *CoreHub) AgentNew(tenantId string, data *entities.Agent) error {
	return c.coredb.AgentNew(tenantId, data)
}

func (c *CoreHub) AgentUpdate(tenantId, pid, id string, data map[string]interface{}) error {
	return c.coredb.AgentUpdate(tenantId, pid, id, data)
}

func (c *CoreHub) AgentGet(tenantId, pid, id string) (*entities.Agent, error) {
	return c.coredb.AgentGet(tenantId, pid, id)
}

func (c *CoreHub) AgentDel(tenantId, pid, agentId string) error {
	return c.coredb.AgentDel(tenantId, pid, agentId)
}

func (c *CoreHub) AgentList(tenantId, pid string) ([]*entities.Agent, error) {
	return c.coredb.AgentList(tenantId, pid)
}

// resource
func (c *CoreHub) ResourceNew(tenantId string, obj *entities.Resource) error {
	return c.coredb.ResourceNew(tenantId, obj)
}

func (c *CoreHub) ResourceUpdate(tenantId string, id string, data map[string]interface{}) error {
	return c.coredb.ResourceUpdate(tenantId, id, data)
}

func (c *CoreHub) ResourceGet(tenantId, rid string) (*entities.Resource, error) {
	return c.coredb.ResourceGet(tenantId, rid)
}

func (c *CoreHub) ResourceDel(tenantId, rid string) error {
	return c.coredb.ResourceDel(tenantId, rid)
}

func (c *CoreHub) ResourceList(tenantId string) ([]*entities.Resource, error) {
	return c.coredb.ResourceList(tenantId)
}

func (c *CoreHub) ResourcesMulti(tenantId string, rids ...string) ([]*entities.Resource, error) {
	return c.coredb.ResourcesMulti(tenantId, rids...)
}

func (c *CoreHub) ResourcesByTarget(tenantId string, target string) ([]*entities.Resource, error) {
	return c.coredb.ResourcesByTarget(tenantId, target)
}

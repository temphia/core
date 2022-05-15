package admin

import (
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/models/vmodels"
)

// plug
func (c *Controller) PlugNew(uclaim *claim.Session, pg *entities.Plug) error {
	err := c.canPlugAction(uclaim, ActionPlugUpdate, pg.Id)
	if err != nil {
		return err
	}

	return c.coredb.PlugNew(uclaim.TenentID, pg)
}

func (c *Controller) PlugUpdate(uclaim *claim.Session, pid string, data map[string]interface{}) error {
	return c.coredb.PlugUpdate(uclaim.TenentID, pid, data)
}

func (c *Controller) PlugGet(uclaim *claim.Session, pid string) (*entities.Plug, error) {
	err := c.canPlugAction(uclaim, ActionPlugGet, pid)
	if err != nil {
		return nil, err
	}

	return c.coredb.PlugGet(uclaim.TenentID, pid)
}

func (c *Controller) PlugDel(uclaim *claim.Session, pid string) error {
	err := c.canPlugAction(uclaim, ActionPlugDel, pid)
	if err != nil {
		return err
	}

	return c.coredb.PlugDel(uclaim.TenentID, pid)
}

func (c *Controller) PlugList(uclaim *claim.Session) ([]*entities.Plug, error) {
	// fixme

	return c.coredb.PlugList(uclaim.TenentID)
}

// agent

func (c *Controller) AgentNew(uclaim *claim.Session, data *entities.Agent) error {
	err := c.canPlugAction(uclaim, ActionAgentUpdate, data.PlugID, data.Id)
	if err != nil {
		return err
	}

	return c.coredb.AgentNew(uclaim.TenentID, data)
}

func (c *Controller) AgentUpdate(uclaim *claim.Session, pid string, aid string, data map[string]interface{}) error {
	return c.coredb.AgentUpdate(uclaim.TenentID, pid, aid, data)
}

func (c *Controller) AgentGet(uclaim *claim.Session, pid, agentId string) (*entities.Agent, error) {
	err := c.canPlugAction(uclaim, ActionAgentUpdate, pid, agentId)
	if err != nil {
		return nil, err
	}

	return c.coredb.AgentGet(uclaim.TenentID, pid, agentId)
}

func (c *Controller) AgentDel(uclaim *claim.Session, pid, agentId string) error {
	err := c.canPlugAction(uclaim, ActionAgentUpdate, pid, agentId)
	if err != nil {
		return err
	}

	return c.coredb.AgentDel(uclaim.TenentID, pid, agentId)
}

func (c *Controller) AgentList(uclaim *claim.Session, pid string) ([]*entities.Agent, error) {
	// fixme =>

	return c.coredb.AgentList(uclaim.TenentID, pid)
}

// resource

func (c *Controller) ResourceNew(uclaim *claim.Session, data *entities.Resource) error {
	return c.coredb.ResourceNew(uclaim.TenentID, data)
}

func (c *Controller) ResourceUpdate(uclaim *claim.Session, rid string, data map[string]interface{}) error {
	return c.coredb.ResourceUpdate(uclaim.TenentID, rid, data)
}

func (c *Controller) ResourceGet(uclaim *claim.Session, rid string) (*entities.Resource, error) {
	return c.coredb.ResourceGet(uclaim.TenentID, rid)
}

func (c *Controller) ResourceDel(uclaim *claim.Session, rid string) error {
	return c.coredb.ResourceDel(uclaim.TenentID, rid)
}

func (c *Controller) ResourceList(uclaim *claim.Session) ([]*entities.Resource, error) {
	return c.coredb.ResourceList(uclaim.TenentID)
}

func (c *Controller) ResourceAgentList(uclaim *claim.Session, req *vmodels.ResourceQuery) ([]*entities.Resource, error) {
	agent, err := c.coredb.AgentGet(uclaim.TenentID, req.PlugId, req.AgentId)
	if err != nil {
		return nil, err
	}

	vals := make([]string, 0, len(agent.Resources))
	for _, v := range agent.Resources {
		vals = append(vals, v)
	}

	resources, err := c.coredb.ResourcesMulti(uclaim.TenentID, vals...)
	return resources, err
}

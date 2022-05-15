package admin

import (
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
)

const (
	ActionPlugUpdate     = "plug_update"
	ActionPlugGet        = "plug_get"
	ActionPlugDel        = "plug_del"
	ActionPlugList       = "plug_list"
	ActionAgentUpdate    = "agent_update"
	ActionAgentGet       = "agent_get"
	ActionAgentDel       = "agent_del"
	ActionAgentList      = "agent_list"
	ActionSignalUpdate   = "agent_signal_update"
	ActionSignalGet      = "agent_signal_get"
	ActionSignalDel      = "agent_signal_del"
	ActionSignalList     = "agent_signal_list"
	ActionResourceUpdate = "agent_resource_update"
	ActionResourceGet    = "agent_resource_get"
	ActionResourceDel    = "agent_resource_del"
	ActionResourceList   = "agent_resource_list"
)

func (c *Controller) canPlugAction(uclaim *claim.Session, action string, path ...string) error {
	if uclaim.IsSuperAdmin() {
		return nil
	}

	err := c.hasPlugScope(uclaim)
	if err != nil {
		return err
	}

	return c.hasPlugPerm(uclaim)
}

func (c *Controller) hasPlugScope(uclaim *claim.Session) error {

	return nil
}

func (c *Controller) hasPlugPerm(uclaim *claim.Session) error {
	// if !uclaim.JitRBAC {
	// 	return nil
	// }

	return easyerr.NotAuthorized()
}

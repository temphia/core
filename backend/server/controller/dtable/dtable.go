package dtable

import (
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/btypes/store"
)

type Controller struct {
	fencer service.Fencer
	dynHub store.DynHub

	cabHub store.CabinetHub
	signer service.Signer
}

func New(fencer service.Fencer, dhub store.DynHub, cabHub store.CabinetHub, signer service.Signer) *Controller {
	return &Controller{
		fencer: fencer,
		dynHub: dhub,
		cabHub: cabHub,
		signer: signer,
	}
}

func (c *Controller) ListSources(uclaim *claim.Session) ([]string, error) {
	return c.dynHub.ListSources(uclaim.TenentID)
}

// dyn_table_group
func (c *Controller) NewGroup(uclaim *claim.Session, source string, model *entities.NewTableGroup) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentID)

	return dynDB.NewGroup(model)
}

func (c *Controller) EditGroup(uclaim *claim.Session, source, gslug string, model *entities.TableGroupPartial) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentID)

	return dynDB.EditGroup(gslug, model)
}

func (c *Controller) GetGroup(uclaim *claim.Session, source, gslug string) (*entities.TableGroup, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentID)
	return dynDB.GetGroup(gslug)
}

func (c *Controller) ListGroup(uclaim *claim.Session, source string) ([]*entities.TableGroup, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentID)
	return dynDB.ListGroup()
}

func (c *Controller) DeleteGroup(uclaim *claim.Session, source, gslug string) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentID)
	return dynDB.DeleteGroup(gslug)
}

// dyn_table
func (c *Controller) AddTable(uclaim *claim.Session, model *entities.NewTable) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.AddTable(uclaim.ServicePath[2], model)
}

func (c *Controller) EditTable(uclaim *claim.Session, tslug string, model *entities.TablePartial) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.EditTable(uclaim.ServicePath[2], tslug, model)
}

func (c *Controller) GetTable(uclaim *claim.Session, tslug string) (*entities.Table, error) {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.GetTable(uclaim.ServicePath[2], tslug)
}

func (c *Controller) ListTables(uclaim *claim.Session) ([]*entities.Table, error) {

	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.ListTables(uclaim.ServicePath[2])
}

func (c *Controller) LoadGroup(uclaim *claim.Session) (*store.LoadDgroupResp, error) {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	tg, err := dynDB.GetGroup(uclaim.ServicePath[2])
	if err != nil {
		return nil, err
	}

	tables, err := dynDB.ListTables(uclaim.ServicePath[2])

	if err != nil {
		return nil, err
	}

	if tg.CabinetSource == "" || tg.CabinetFolder == "" {
		tg.CabinetSource = c.cabHub.DefaultName(uclaim.TenentID)
		tg.CabinetFolder = "data_common"
	}

	cabToken, err := c.signer.Sign(uclaim.TenentID, &claim.TicketCabinet{
		Mode:     "",
		Folder:   tg.CabinetFolder,
		Source:   tg.CabinetSource,
		Expiry:   0,
		Prefix:   "",
		DeviceId: uclaim.DeviceID,
	})
	if err != nil {
		return nil, err
	}

	resp := &store.LoadDgroupResp{
		Tables:          tables,
		CabinetTicket:   cabToken,
		SockdRoomTicket: "",
	}

	return resp, nil
}

func (c *Controller) DeleteTable(uclaim *claim.Session, tslug string) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.DeleteTable(uclaim.ServicePath[2], tslug)
}

// dyn_table_column
func (c *Controller) AddColumn(uclaim *claim.Session, tslug string, model *entities.NewColumn) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.AddColumn(uclaim.ServicePath[2], tslug, model)
}

func (c *Controller) GetColumn(uclaim *claim.Session, tslug, cslug string) (*entities.Column, error) {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.GetColumn(uclaim.ServicePath[2], tslug, cslug)
}

func (c *Controller) EditColumn(uclaim *claim.Session, tslug, cslug string, model *entities.ColumnPartial) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return dynDB.EditColumn(uclaim.ServicePath[2], tslug, cslug, model)
}

func (c *Controller) ListColumns(uclaim *claim.Session, tslug string) ([]*entities.Column, error) {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return dynDB.ListColumns(uclaim.ServicePath[2], tslug)
}

func (c *Controller) DeleteColumn(uclaim *claim.Session, tslug, cslug string) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return dynDB.DeleteColumn(uclaim.ServicePath[2], tslug, cslug)
}

func (c *Controller) AddIndex(uclaim *claim.Session, tslug string, model *entities.Index) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.AddIndex(uclaim.ServicePath[2], tslug, model)
}

// dyn_table_meta
func (c *Controller) AddUniqueIndex(uclaim *claim.Session, tslug string, model *entities.Index) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.AddUniqueIndex(uclaim.ServicePath[2], tslug, model)
}

func (c *Controller) AddFTSIndex(uclaim *claim.Session, tslug string, model *entities.FTSIndex) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.AddFTSIndex(uclaim.ServicePath[2], tslug, model)
}

func (c *Controller) AddColumnFRef(uclaim *claim.Session, tslug string, model *entities.ColumnFKRef) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.AddColumnFRef(uclaim.ServicePath[2], tslug, model)
}

func (c *Controller) ListIndex(uclaim *claim.Session, tslug string) ([]*entities.Index, error) {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.ListIndex(uclaim.ServicePath[2], tslug)
}

func (c *Controller) RemoveIndex(uclaim *claim.Session, tslug, slug string) error {
	dynDB := c.dynHub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return dynDB.RemoveIndex(uclaim.ServicePath[2], tslug, slug)
}

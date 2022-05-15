package cabinet

import (
	"context"

	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/btypes/store"
)

type Controller struct {
	fencer service.Fencer
	hub    store.CabinetHub
	signer service.Signer
}

func New(fencer service.Fencer, cabinet store.CabinetHub, signer service.Signer) *Controller {
	return &Controller{
		fencer: fencer,
		hub:    cabinet,
		signer: signer,
	}
}

func (c *Controller) ListRoot(uclaim *claim.Session) ([]string, error) {
	if uclaim.ServicePath[0] != "cabinet" {
		return nil, easyerr.NotAuthorized()
	}

	sourced := c.hub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return sourced.ListRoot(context.TODO())
}

func (c *Controller) AddFolder(uclaim *claim.Session, folder string) error {
	err := c.canAction(uclaim, "add_folder", folder)
	if err != nil {
		return err
	}
	sourced := c.hub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)
	return sourced.AddFolder(context.TODO(), folder)
}

func (c *Controller) AddBlob(uclaim *claim.Session, folder, file string, contents []byte) error {
	err := c.canAction(uclaim, "add_blob", (folder))
	if err != nil {
		return err
	}
	sourced := c.hub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return sourced.AddBlob(context.TODO(), folder, file, contents)
}

func (c *Controller) ListFolder(uclaim *claim.Session, folder string) ([]*store.BlobInfo, error) {
	err := c.canAction(uclaim, "list_folder", folder)
	if err != nil {
		return nil, err
	}

	sourced := c.hub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return sourced.ListFolder(context.TODO(), folder)
}

func (c *Controller) GetBlob(uclaim *claim.Session, folder, file string) ([]byte, error) {
	err := c.canAction(uclaim, "get_blob", folder)
	if err != nil {
		return nil, err
	}
	sourced := c.hub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return sourced.GetBlob(context.TODO(), folder, file)
}

func (c *Controller) DeleteBlob(uclaim *claim.Session, folder, file string) error {
	err := c.canAction(uclaim, "del_blob", folder)
	if err != nil {
		return err
	}
	sourced := c.hub.GetSource(uclaim.ServicePath[1], uclaim.TenentID)

	return sourced.DeleteBlob(context.TODO(), folder, file)
}

func (c *Controller) NewFolderTicket(uclaim *claim.Session, folder string) (string, error) {
	return c.signer.Sign(uclaim.TenentID, &claim.TicketCabinet{
		Mode:     "",
		Folder:   folder,
		Source:   uclaim.ServicePath[1],
		Expiry:   0,
		Prefix:   "",
		DeviceId: uclaim.DeviceID,
	})
}

// Ticket cabinet
func (c *Controller) TicketFile(tenantId, file string, ticket *claim.TicketCabinet) ([]byte, error) {
	// if !ticket.AllowGet {
	// 	return nil, easyerr.NotAuthorized()
	// }

	// if ticket.PinnedFiles != nil {
	// 	if !funk.ContainsString(ticket.PinnedFiles, file) {
	// 		return nil, easyerr.NotAuthorized()
	// 	}
	// }

	// fixme => check prefix

	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.GetBlob(context.TODO(), ticket.Folder, file)
}

func (c *Controller) TicketPreview(tenantId, file string, ticket *claim.TicketCabinet) ([]byte, error) {
	// if !ticket.AllowGet {
	// 	return nil, easyerr.NotAuthorized()
	// }
	// fixme => implement preview
	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.GetBlob(context.TODO(), ticket.Folder, file)
}

func (c *Controller) TicketList(tenantId string, ticket *claim.TicketCabinet) ([]*store.BlobInfo, error) {
	// pp.Println(ticket)
	// if !ticket.AllowList {
	// 	return nil, easyerr.NotAuthorized()
	// }

	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.ListFolder(context.TODO(), ticket.Folder)
}

func (c *Controller) TicketUpload(tenantId, file string, data []byte, ticket *claim.TicketCabinet) error {
	// fixme =>  send back upload proof token
	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.AddBlob(context.TODO(), ticket.Folder, file, data)
}

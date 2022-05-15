package pacman

import (
	"context"
	"encoding/json"

	"github.com/rs/xid"
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/models/vmodels"
	"github.com/thoas/go-funk"
)

func (c *PacMan) BprintList(tenantid, group string) ([]*entities.BPrint, error) {
	return c.syncer.BprintList(tenantid, group)
}

func (c *PacMan) BprintCreate(tenantid string, bp *entities.BPrint) (string, error) {
	if bp.ID == "" {
		bp.ID = xid.New().String()
	}

	return bp.ID, c.syncer.BprintNew(tenantid, bp)
}

func (c *PacMan) BprintUpdate(tenantid string, bp *entities.BPrint) error {
	if bp.ID == "" || bp.Slug == "" {
		return easyerr.NotFound()
	}
	//_, err := c.localStore.BprintUpdate(tenantid, bp)

	return easyerr.NotImpl()
}

func (c *PacMan) BprintGet(tenantid, bid string) (*entities.BPrint, error) {
	return c.syncer.BprintGet(tenantid, bid)

}

func (c *PacMan) BprintRemove(tenantid, bid string) error {
	bprint, err := c.syncer.BprintGet(tenantid, bid)
	if err != nil {
		return err
	}
	for _, file := range bprint.Files {
		err = c.blobStore(tenantid).DeleteBlob(context.TODO(), btypes.BprintBlobFolder, ffile(bid, file))
	}
	return c.syncer.BprintDel(tenantid, bid)
}

func (c *PacMan) BprintListBlobs(tenantid, bid string) (interface{}, error) {
	return nil, nil
}

func (c *PacMan) BprintNewBlob(tenantid, bid, file string, payload []byte) error {
	bprint, err := c.BprintGet(tenantid, bid)
	if err != nil {
		return err
	}

	err = c.blobStore(tenantid).AddBlob(context.TODO(), btypes.BprintBlobFolder, ffile(bid, file), payload)
	if err != nil {
		return err
	}

	if funk.ContainsString(bprint.Files, file) {
		return nil
	}

	bprint.Files = append(bprint.Files, file)
	bprint.ID = bid

	err = c.syncer.BprintUpdate(tenantid, bid, map[string]interface{}{
		"files": bprint.Files,
	})

	if err != nil {
		c.BprintDeleteBlob(tenantid, bid, file)
		return easyerr.Error("could not finish blob add")
	}
	return nil
}

func (c *PacMan) BprintUpdateBlob(tenantid, bid, file string, payload []byte) error {
	return c.blobStore(tenantid).AddBlob(context.TODO(), btypes.BprintBlobFolder, ffile(bid, file), payload)
}

func (c *PacMan) BprintGetBlob(tenantid, bid, file string) ([]byte, error) {
	return c.blobStore(tenantid).GetBlob(context.TODO(), btypes.BprintBlobFolder, ffile(bid, file))
}
func (c *PacMan) BprintDeleteBlob(tenantid, bid, file string) error {
	// fixme => also delete from list in  bprint
	return c.blobStore(tenantid).DeleteBlob(context.TODO(), ffile(bid, file), file)
}

func (b *PacMan) Install(tenantId string, opts *vmodels.RepoInstallOpts) (interface{}, error) {
	bprint, err := b.syncer.BprintGet(tenantId, opts.BprintId)
	if err != nil {
		return nil, err
	}

	switch bprint.Type {
	case btypes.BPrintTypePlug, btypes.BprintTypePlugin:
		popts := &vmodels.PlugInstallOptions{}
		err := json.Unmarshal(opts.Data, popts)
		if err != nil {
			return nil, err
		}
		return b.installPlug(tenantId, bprint, popts)
	case btypes.BPrintTypeTSchema:
		popts := &vmodels.DGroupInstallOptions{}
		err := json.Unmarshal(opts.Data, popts)
		if err != nil {
			return nil, err
		}

		popts.UserId = opts.UserId
		return nil, b.installDGroup(tenantId, bprint, popts)

	default:
		return nil, easyerr.NotImpl()
	}

}

func ffile(id, file string) string {
	return id + "_" + file
}

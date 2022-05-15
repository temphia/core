package dynddl

import (
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/store/upper/dyndb/dyncore"
	"github.com/temphia/core/backend/server/store/upper/ucore"
	"github.com/upper/db/v4"
)

type DynDDL struct {
	session    db.Session
	sharedLock service.DyndbLock
	dyngen     ucore.Zenerator
}

func New(session db.Session, sharedLock service.DyndbLock, dyngen ucore.Zenerator) *DynDDL {
	return &DynDDL{
		session:    session,
		sharedLock: sharedLock,
		dyngen:     dyngen,
	}
}

func (d *DynDDL) dataTableGroups() db.Collection {
	return dyncore.GroupTable(d.session)
}

func (d *DynDDL) dataTables() db.Collection {
	return dyncore.Table(d.session)
}

func (d *DynDDL) dataTableColumns() db.Collection {
	return dyncore.TableColumn(d.session)
}

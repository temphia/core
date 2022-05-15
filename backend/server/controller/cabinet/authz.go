package cabinet

import (
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/thoas/go-funk"
)

func (c *Controller) canAction(uclaim *claim.Session, target, action string) error {
	if uclaim.IsSuperAdmin() {
		return nil
	}

	err := c.scopeCheck(uclaim, target, action)
	if err != nil {
		return err
	}

	// if uclaim.SkipRBAC {
	// 	return nil
	// }

	return c.permCheck(uclaim, target, action)
}

func (c *Controller) scopeCheck(uclaim *claim.Session, target, action string) error {
	if !(uclaim.ServicePath[0] == "cabinet") {
		return easyerr.NotAuthorized()
	}

	for _, object := range uclaim.Objects {
		objPath := object.Path[0]

		if objPath == target || objPath == "*" {
			if object.WhitelistMode {
				if funk.ContainsString(object.Actions, action) {
					return nil
				}
				return easyerr.NotAuthorized()
			}

			if funk.ContainsString(object.Actions, action) {
				// blacklisted action
				return easyerr.NotAuthorized()
			}
			return nil
		}
	}

	return easyerr.NotAuthorized()
}

func (c *Controller) permCheck(uclaim *claim.Session, target, action string) error {
	return easyerr.NotImpl()
}

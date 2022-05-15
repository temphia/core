package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/vmodels"
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/temphia/temphia/backend/server/btypes/store"
	"github.com/temphia/temphia/backend/server/lib/apiutils"
)

type Controller struct {
	coredb store.CoreHub
	fencer service.Fencer
	signer service.Signer
}

func New(coredb store.CoreHub, fencer service.Fencer, signer service.Signer) *Controller {
	return &Controller{
		coredb: coredb,
		fencer: fencer,
		signer: signer,
	}
}

func (c *Controller) LogInUser(tenantId, userOrEmail, password string, ctx *gin.Context) {
	err := c.logInUser(tenantId, userOrEmail, password, ctx)
	if err == nil {
		return
	}
	apiutils.WriteErr(ctx, err.Error())
}

func (c *Controller) RefreshService(uclaim *claim.User, opts *vmodels.RefreshReq) *vmodels.RefreshResp {
	return c.refreshService(uclaim, opts)
}

/*
	// poc

	func (c *Controller) SignUpToken() {}
	func (c *Controller) SignUpDirect() {}

*/

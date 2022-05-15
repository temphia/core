package auth

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/vmodels"
)

func (c *Controller) refreshService(uclaim *claim.User, opts *vmodels.RefreshReq) *vmodels.RefreshResp {

	if uclaim.IsSuperAdmin() {
		return c.superClaim(uclaim, opts)
	}

	switch opts.Path[0] {
	case "dgroup":
		return c.dgroupClaim(uclaim, opts)
	case "cabinet":
		return c.CabinetClaim(uclaim, opts)
	case "exector":
		return c.execClaim(uclaim, opts)
	case "admin":
		return c.adminClaim(uclaim, opts)
	default:
		panic("not supported")
	}
}

func (c *Controller) superClaim(uclaim *claim.User, opts *vmodels.RefreshReq) *vmodels.RefreshResp {
	deviceId := xid.New().String()
	serviceId := ""
	if opts.OldToken != "" {
		sess := &claim.Session{}
		err := c.signer.Parse(uclaim.TenentID, opts.OldToken, sess)
		if err != nil {
			deviceId = sess.DeviceID
		}
	}

	token, err := c.signer.Sign(uclaim.TenentID, &claim.Session{
		TenentID:    uclaim.TenentID,
		UserID:      uclaim.UserID,
		UserGroup:   uclaim.UserGroup,
		Type:        "session",
		Expiry:      0,
		Attributes:  nil,
		SessionID:   serviceId,
		DeviceID:    deviceId,
		ServicePath: opts.Path,
		RBACMode:    claim.RBACSkip,
		Objects:     nil,
	})

	if err != nil {
		return &vmodels.RefreshResp{
			Token:    "",
			Expiry:   "",
			Message:  err.Error(),
			StatusOk: false,
		}
	}

	return &vmodels.RefreshResp{
		Token:    token,
		Expiry:   "",
		Message:  "",
		StatusOk: true,
	}

}

func (c *Controller) dgroupClaim(uclaim *claim.User, opts *vmodels.RefreshReq) *vmodels.RefreshResp {
	return &vmodels.RefreshResp{
		Token:    "",
		Expiry:   "",
		Message:  "",
		StatusOk: false,
	}
}

func (c *Controller) CabinetClaim(uclaim *claim.User, opts *vmodels.RefreshReq) *vmodels.RefreshResp {
	return &vmodels.RefreshResp{
		Token:    "",
		Expiry:   "",
		Message:  "",
		StatusOk: false,
	}
}

func (c *Controller) execClaim(uclaim *claim.User, opts *vmodels.RefreshReq) *vmodels.RefreshResp {
	return &vmodels.RefreshResp{
		Token:    "",
		Expiry:   "",
		Message:  "",
		StatusOk: false,
	}
}

func (c *Controller) adminClaim(uclaim *claim.User, opts *vmodels.RefreshReq) *vmodels.RefreshResp {

	return &vmodels.RefreshResp{
		Token:    "",
		Expiry:   "",
		Message:  "",
		StatusOk: false,
	}
}

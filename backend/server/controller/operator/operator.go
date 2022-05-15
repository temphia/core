package operator

import (
	"encoding/json"
	"time"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"

	"github.com/temphia/core/backend/server/app/config"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/temphia/core/backend/server/btypes/models/vmodels"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/temphia/core/backend/server/btypes/store"
)

type Controller struct {
	coredb     store.CoreHub // fixme => use control plane instead of coredb sirectly
	signer     service.Signer
	OpUser     string
	OpPassword string
	OpToken    string
}

func New(cdb store.CoreHub, signer service.Signer, config *config.AppConfig) *Controller {

	return &Controller{
		signer:     signer,
		coredb:     cdb,
		OpUser:     config.OperatorName,
		OpPassword: config.OperatorPassword,
		OpToken:    "",
	}
}

func (c *Controller) Login(data *vmodels.OperatorLoginReq) (*vmodels.OperatorLoginResp, error) {
	pp.Println("@@=>", c.OpUser, c.OpPassword, c.OpToken, data)

	if data.MasterOpToken != "" {
		if !c.verifyMasterToken() {
			return nil, easyerr.Error("Invalid Mastertoken")
		}
	} else {
		// fixme => security and stuff (constant time compare ?)
		if data.User != c.OpUser || data.Password != c.OpPassword {
			return nil, easyerr.Error("Invalid User crediantials")
		}
	}

	opClaim := &claim.OperatorClaim{
		XID:          xid.New().String(),
		ClaimType:    "operator",
		Expiry:       0,
		Origin:       "",
		BindDeviceId: "",
	}

	ocBytes, err := json.Marshal(opClaim)
	if err != nil {
		return nil, err
	}
	token, err := c.signer.GlobalSignRaw(string(ocBytes))
	if err != nil {
		return nil, err
	}

	return &vmodels.OperatorLoginResp{
		Token: token,
	}, nil

}

func (c *Controller) AddTenant(data *vmodels.NewTenant) error {
	err := c.coredb.AddTenant(&entities.Tenant{
		Name: data.Name,
		Slug: data.Slug,
	})

	if err != nil {
		return err
	}

	err = c.coredb.AddUserGroup(&entities.UserGroup{
		Name:     "Super Admin",
		Slug:     "super_admin",
		Icon:     "",
		TenantID: data.Slug,
	})

	if err != nil {
		return err
	}

	return c.coredb.AddUser(&entities.User{
		UserId:    "superuser",
		FullName:  "Super User",
		Email:     data.SuperEmail,
		GroupID:   "super_admin",
		Password:  data.SuperPassword,
		TenantID:  data.Slug,
		PublicKey: "",
		Data:      nil,
		CreatedAt: time.Now(),
		Active:    true,
	})

}

func (c *Controller) UpdateTenant(slug string, data map[string]interface{}) error {
	return c.coredb.UpdateTenant(slug, data)
}

func (c *Controller) ListTenant() ([]*entities.Tenant, error) {
	return c.coredb.ListTenant()
}

func (c *Controller) DeleteTenant(slug string) error {
	return c.coredb.RemoveTenant(slug)
}

func (c *Controller) Stats() {

}
func (c *Controller) TenantToken() {

}

func (c *Controller) ParseToken(raw string) (*claim.OperatorClaim, error) {
	data, err := c.signer.GlobalParseRaw(raw)
	if err != nil {
		return nil, err
	}

	opc := &claim.OperatorClaim{}

	err = json.Unmarshal([]byte(data), opc)
	if err != nil {
		return nil, err
	}

	return opc, nil

}

// private

func (c *Controller) verifyMasterToken() bool {
	return false
}

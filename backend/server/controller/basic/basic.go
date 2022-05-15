package basic

import (
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/service"
	"github.com/temphia/temphia/backend/server/btypes/store"
)

type Controller struct {
	coredb  store.CoreHub
	cabinet store.CabinetHub
	dynHub  store.DynHub
	pacman  service.Pacman
}

func New(coredb store.CoreHub, cabinet store.CabinetHub, dynHub store.DynHub, pacman service.Pacman) *Controller {
	ctrl := &Controller{
		coredb:  coredb,
		cabinet: cabinet,
		dynHub:  dynHub,
		pacman:  pacman,
	}

	return ctrl
}

func (c *Controller) ListConsoleStats(uclaim *claim.Session) (interface{}, error) {
	return map[string]interface{}{
		"abc": "xyz",
	}, nil
}

func (c *Controller) ListRepoSources(uclaim *claim.Session) (interface{}, error) {
	return c.pacman.RepoSourceList(uclaim.TenentID)
}

func (c *Controller) ListCabinetSources(uclaim *claim.Session) ([]string, error) {
	return c.cabinet.ListSources(uclaim.TenentID)
}

func (c *Controller) ListDyndbSources(uclaim *claim.Session) ([]string, error) {
	return c.dynHub.ListSources((uclaim.TenentID))
}

func (c *Controller) JoinNotification() error {
	return nil
}

func (c *Controller) MessageUser(uclaim *claim.Session, userId, message string) (int64, error) {

	return c.coredb.AddUserMessage(&entities.UserMessage{
		Id:           0,
		Title:        "message",
		Read:         false,
		Type:         "user_message",
		Contents:     message,
		UserId:       userId,
		FromUser:     uclaim.UserID,
		FromPlug:     "",
		FromAgent:    "",
		PlugCallback: "",
		WarnLevel:    1,
		Encrypted:    false,
		CreatedAt:    nil,
		TenantId:     uclaim.TenentID,
	})
}

func (c *Controller) GetUserInfo(uclaim *claim.Session, userId string) (*entities.UserInfo, error) {
	usr, err := c.coredb.GetUserByID(uclaim.TenentID, userId)
	if err != nil {
		return nil, err
	}

	fuser := &entities.UserInfo{
		UserId:    uclaim.UserID,
		FullName:  usr.FullName,
		Bio:       "",
		PublicKey: "",
		Email:     "",
		Group:     "",
	}

	return fuser, nil
}

func (c *Controller) GetSelfInfo(uclaim *claim.Session) (*entities.UserInfo, error) {
	usr, err := c.coredb.GetUserByID(uclaim.TenentID, uclaim.UserID)
	if err != nil {
		return nil, err
	}

	fuser := &entities.UserInfo{
		UserId:    uclaim.UserID,
		FullName:  usr.FullName,
		PublicKey: usr.PublicKey,
		Bio:       "",
		Email:     usr.Email,
		Group:     uclaim.UserGroup,
	}

	return fuser, nil
}

func (c *Controller) GetChangeEmail(uclaim *claim.Session) error {

	return nil
}

func (c *Controller) GetChangeAuth(uclaim *claim.Session) error {
	return nil
}

func (c *Controller) ListMessages(uclaim *claim.Session, opts *entities.UserMessageReq) ([]*entities.UserMessage, error) {
	opts.UserId = uclaim.UserID
	return c.coredb.ListUserMessages(uclaim.TenentID, opts)
}

func (c *Controller) ModifyMessages(uclaim *claim.Session, opts *entities.ModifyMessages) error {
	switch opts.Operation {
	case "delete":
		return c.coredb.DeleteUserMessages(uclaim.TenentID, uclaim.UserID, opts.Ids)
	case "read":
		return c.coredb.ReadUserMessages(uclaim.TenentID, uclaim.UserID, opts.Ids)
	default:
		panic("not impl")
	}
}

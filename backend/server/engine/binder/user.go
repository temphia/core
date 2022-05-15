package binder

import "github.com/temphia/temphia/backend/server/btypes/models/entities"

func (b *Binder) ListUsers(group string) ([]string, error) {
	users, err := b.factory.App.CoreHub().ListUsersByGroup(b.namespace, group)
	if err != nil {
		return nil, err
	}

	ustrs := make([]string, 0, len(users))

	for _, u := range users {
		ustrs = append(ustrs, u.UserId)
	}

	return ustrs, nil
}

func (b *Binder) MessageUser(group, user, message string, encrypted bool) error {

	_, err := b.factory.App.CoreHub().AddUserMessage(&entities.UserMessage{
		Id:           0,
		Title:        "plug message",
		Read:         false,
		Type:         "message",
		Contents:     message,
		UserId:       user,
		FromUser:     "",
		FromPlug:     b.plugId,
		FromAgent:    b.agentId,
		PlugCallback: "",
		WarnLevel:    1,
		Encrypted:    false,
		CreatedAt:    nil,
		TenantId:     b.namespace,
	})

	return err
}

func (b *Binder) GetUser(group, name string) (*entities.UserInfo, error) {
	user, err := b.factory.App.CoreHub().GetUserByID(b.namespace, name)
	if err != nil {
		return nil, err
	}

	usr := &entities.UserInfo{
		UserId:    user.UserId,
		FullName:  user.FullName,
		Bio:       user.Bio,
		PublicKey: user.PublicKey,
		Email:     user.Email,
		Group:     user.GroupID,
	}

	return usr, nil
}

// current user

func (b *Binder) MessageCurrentUser(user, message string, encrypted bool) error {
	return nil
}

func (b *Binder) CurrentUser() (*entities.UserInfo, error) {
	return nil, nil
}

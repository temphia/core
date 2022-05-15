package coredb

import (
	"github.com/temphia/core/backend/server/btypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) AddUser(user *entities.User) error {
	_, err := d.userTable().Insert(user)
	return err
}

func (d *DB) UpdateUser(tenantId, user string, data map[string]interface{}) error {
	return d.userTable().Find("tenant_id", tenantId, "user_id", user).Update(data)
}

func (d *DB) RemoveUser(tenantId string, username string) error {
	return d.userTable().Find(db.Cond{"tenant_id": tenantId, "user_id": username}).Delete()
}

func (d *DB) GetUserByID(tenantId string, username string) (*entities.User, error) {
	usr := &entities.User{}
	err := d.userTable().Find(
		db.Cond{
			"tenant_id": tenantId,
			"user_id":   username,
		},
	).One(usr)
	return usr, err
}

func (d *DB) GetUserByEmail(tenantId string, email string) (*entities.User, error) {
	usr := &entities.User{}
	err := d.userTable().Find(
		db.Cond{
			"tenant_id": tenantId,
			"email":     email,
		},
	).One(usr)

	return usr, err
}

func (d *DB) ListUsers(tenantId string) ([]*entities.User, error) {
	us := make([]*entities.User, 0, 10)
	err := d.userTable().Find(
		db.Cond{
			"tenant_id": tenantId,
		},
	).All(&us)

	return us, err
}

func (d *DB) ListUsersByGroup(tenantId string, groupid string) ([]*entities.User, error) {
	us := make([]*entities.User, 0)

	cond := db.Cond{
		"tenant_id": tenantId,
		"group_id":  groupid,
	}

	err := d.userTable().Find(cond).All(&us)
	return us, err
}

func (d *DB) ListUserWithTags(tenantId string, tags ...string) ([]*entities.User, error) {
	return nil, nil
}

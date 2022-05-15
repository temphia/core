package dbdata

import (
	"github.com/temphia/temphia/backend/server/btypes/store"
)

func DevUnsafeSeed(d store.CoreDB) error {

	// err := d.AddTenant(&entities.Tenant{
	// 	Name:   "Shell Corp",
	// 	Slug:   "ten1",
	// 	Config: "",
	// })

	// if err != nil {
	// 	return err
	// }

	// err = d.AddUserGroup(&entities.UserGroup{
	// 	Slug:        "dev",
	// 	TenantID:    "ten1",
	// 	LoginPolicy: "",
	// })

	// err = d.AddUser(&entities.User{
	// 	Email:                 "jd@jd.com",
	// 	GroupID:               "dev",
	// 	Firstname:             "Jane",
	// 	Lastname:              "doe",
	// 	Password:              "xyz123",
	// 	PendingPasswordChange: true,
	// 	Admin:                 true,
	// 	TenantID:              "ten1",
	// 	Username:              "jd1",
	// 	AuthProvider:          "",
	// 	Config:                "",
	// 	CreatedAt:             "",
	// })

	// if err != nil {
	// 	return err
	// }

	return nil
}

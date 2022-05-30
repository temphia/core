package store

import (
	"github.com/temphia/core/backend/server/btypes/models/entities"
)

type CoreHub interface {
	UserOps
	TenantOps
	SyncDB
	UserMessageOps
	UserGroupExtra
}

type SyncDB interface {
	BprintOps
	PlugOps
}

type CoreDB interface {
	CoreHub

	Migrate() error
	GetInnerDriver() interface{}
}

type TenantOps interface {
	AddTenant(tenant *entities.Tenant) error
	UpdateTenant(slug string, data map[string]interface{}) error
	GetTenant(tenant string) (*entities.Tenant, error)
	RemoveTenant(slug string) error
	ListTenant() ([]*entities.Tenant, error)

	AddDomain(domain *entities.TenantDomain) error
	UpdateDomain(tenantId string, id int64, data map[string]interface{}) error
	GetDomain(tenantId string, id int64) (*entities.TenantDomain, error)
	RemoveDomain(tenantId string, id int64) error
	ListDomain(tenantId string) ([]*entities.TenantDomain, error)

	AddDomainWidget(domain *entities.DomainWidget) error
	UpdateDomainWidget(tenantId string, id int64, data map[string]interface{}) error
	GetDomainWidget(tenantId string, id int64) (*entities.DomainWidget, error)
	RemoveDomainWidget(tenantId string, id int64) error
	ListDomainWidget(tenantId string, did int64) ([]*entities.DomainWidget, error)
}

type UserOps interface {
	AddUserGroup(ug *entities.UserGroup) error
	GetUserGroup(tenantId string, slug string) (*entities.UserGroup, error)
	ListUserGroups(tenantId string) ([]*entities.UserGroup, error)
	UpdateUserGroup(tenantId, slug string, data map[string]interface{}) error
	RemoveUserGroup(tenantId string, ugslug string) error

	AddUser(user *entities.User) error
	UpdateUser(tenantId, user string, data map[string]interface{}) error
	RemoveUser(tenantId string, username string) error
	GetUserByID(tenantId string, username string) (*entities.User, error)
	GetUserByEmail(tenantId string, email string) (*entities.User, error)
	ListUsers(tenantId string) ([]*entities.User, error)
	ListUsersByGroup(tenantId, group string) ([]*entities.User, error)
}

type UserPermissionOps interface {
	AddPerm(data *entities.Permission) error
	UpdatePerm(data map[string]interface{}) error
	GetPerm(tenantId string, id int64) (*entities.Permission, error)
	RemovePerm(tenantId string, id int64) error

	AddRole(data *entities.Role) error
	GetRole(tenantId string, id int64) (*entities.Role, error)
	UpdateRole(data map[string]interface{}) error
	RemoveRole(data *entities.Role) error
	AddUserRole(data *entities.UserRole) error
	GetUserRole(tenantId string, id int64) (*entities.UserRole, error)
	UpdateUserRole(data map[string]interface{}) error
	RemoveUserRole(data *entities.UserRole) error
	ListAllPerm(tenantId string) ([]*entities.Permission, error)
	ListAllRole(tenantId string) ([]*entities.Permission, error)
	ListAllUserRole(tenantId string) ([]*entities.Permission, error)
	ListAllUserPerm(tenantId string) ([]*entities.Permission, error)
	ListUserPerm(tenantId string, userId, objType, objsub string) ([]*entities.Permission, error)
}

type UserGroupExtra interface {
	AddUserGroupAuth(data *entities.UserGroupAuth) error
	UpdateUserGroupAuth(tenantId string, gslug string, id int64, data map[string]interface{}) error
	ListUserGroupAuth(tenantId string, gslug string) ([]*entities.UserGroupAuth, error)
	GetUserGroupAuth(tenantId string, gslug string, id int64) (*entities.UserGroupAuth, error)
	RemoveUserGroupAuth(tenantId, gslug string, id int64) error

	AddUserGroupHook(data *entities.UserGroupHook) error
	UpdateUserGroupHook(tenantId string, gslug string, id int64, data map[string]interface{}) error
	ListUserGroupHook(tenantId string, gslug string) ([]*entities.UserGroupHook, error)
	GetUserGroupHook(tenantId string, gslug string, id int64) (*entities.UserGroupHook, error)
	RemoveUserGroupHook(tenantId, gslug string, id int64) error

	AddUserGroupPlug(data *entities.UserGroupPlug) error
	UpdateUserGroupPlug(tenantId string, gslug string, id int64, data map[string]interface{}) error
	ListUserGroupPlug(tenantId string, gslug string) ([]*entities.UserGroupPlug, error)
	GetUserGroupPlug(tenantId string, gslug string, id int64) (*entities.UserGroupPlug, error)
	RemoveUserGroupPlug(tenantId, gslug string, id int64) error

	AddUserGroupData(data *entities.UserGroupData) error
	UpdateUserGroupData(tenantId string, gslug string, id int64, data map[string]interface{}) error
	ListUserGroupData(tenantId string, gslug string) ([]*entities.UserGroupData, error)
	GetUserGroupData(tenantId string, gslug string, id int64) (*entities.UserGroupData, error)
	RemoveUserGroupData(tenantId, gslug string, id int64) error
}

type UserMessageOps interface {
	AddUserMessage(msg *entities.UserMessage) (int64, error)
	UserMessageSetRead(tenantId, user string, id int64) error
	RemoveUserMessage(tenantId string, username string, id int64) error
	ListUserMessages(tenantId string, data *entities.UserMessageReq) ([]*entities.UserMessage, error)

	ReadUserMessages(tenantId, userId string, id []int64) error
	DeleteUserMessages(tenantId, userId string, id []int64) error
}

// sync_db

type BprintOps interface {
	BprintNew(tenantId string, et *entities.BPrint) error
	BprintUpdate(tenantId, id string, data map[string]interface{}) error
	BprintGet(tenantId, id string) (*entities.BPrint, error)
	BprintDel(tenantId, id string) error
	BprintList(tenantId, group string) ([]*entities.BPrint, error)
}

type PlugOps interface {
	PlugNew(tenantId string, pg *entities.Plug) error
	PlugUpdate(tenantId string, id string, data map[string]interface{}) error
	PlugGet(tenantId, pid string) (*entities.Plug, error)
	PlugDel(tenantId, pid string) error
	PlugList(tenantId string) ([]*entities.Plug, error)

	AgentNew(tenantId string, data *entities.Agent) error
	AgentUpdate(tenantId, pid, id string, data map[string]interface{}) error
	AgentGet(tenantId, pid, id string) (*entities.Agent, error)
	AgentDel(tenantId, pid, agentId string) error
	AgentList(tenantId, pid string) ([]*entities.Agent, error)

	ResourceNew(tenantId string, obj *entities.Resource) error
	ResourceUpdate(tenantId string, id string, data map[string]interface{}) error
	ResourceGet(tenantId, rid string) (*entities.Resource, error)
	ResourceDel(tenantId, rid string) error
	ResourceList(tenantId string) ([]*entities.Resource, error)
	ResourcesMulti(tenantId string, rids ...string) ([]*entities.Resource, error)
	ResourcesByTarget(tenantId string, target string) ([]*entities.Resource, error)
}

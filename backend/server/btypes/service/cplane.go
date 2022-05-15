package service

import (
	"github.com/temphia/temphia/backend/server/app/config"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/job"
)

type NodeStats struct {
	Epoch int
	Mem   int
	CPU   int
}

type ClusterStats struct {
	Id    string
	Nodes map[string][]NodeStats
}

type ControlPlane interface {
	Start() error
	Inject(iapp interface{}, config *config.AppConfig) error
	SetJobChan(chan *job.Job)
	NotifyStat(stats NodeStats) error

	GetAppStats() (*ClusterStats, error)
	GetLocker() Locker
	GetRouter() Router

	GetSockdRouter() SockdRouter
	GetEventBus() EventBus
}

type EventBus interface {
	EmitTenantEvent(tenant string, event string, data *entities.Tenant)
	EmitUserGroupEvent(event string, data *entities.UserGroup)
	EmitSchemaChange(tenant, source, group string, data interface{})

	OnDynSchemaChange(fn func(tenant, source, group string, data interface{}))
	OnTenantChange(fn func(tenant string, event string, data *entities.Tenant))
	OnUserGroupChange(fn func(event string, data *entities.UserGroup))
}

type Router interface {
	Route(j *job.Job) bool
}

type SockdRouter interface {
	Publish(tenantId, room string, tags map[string]string, rawData []byte) error
	Broadcast(tenantId, room string, rawData []byte) error
	SendSession(tenantId, session string, rawData []byte) error
}

type Locker interface{}

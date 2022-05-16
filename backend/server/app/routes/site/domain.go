package site

import "github.com/temphia/core/backend/server/btypes/models/entities"

type tenantCache struct {
	domains map[string]*entities.TenantDomain
	widgets []*entities.WidgetHook
}

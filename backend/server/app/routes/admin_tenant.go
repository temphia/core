package routes

import (
	"github.com/temphia/temphia/backend/server/lib/apiutils/request"
)

func (r *R) UpdateTenant(ctx request.Ctx) {
	// ten := &entities.Tenant{}
	// if ten.Slug != cm.TenentID {
	// 	apiutils.WriteErr(c, ErrNotAllowd)
	// 	return
	// }
	// err := r.syncCtrl.UpdateTenant(ten)
	// apiutils.WriteFinal(c, err)
}

// temphia_clusterXXX => xyzmno

func (r *R) NewTenantDomain(ctx request.Ctx)    {}
func (r *R) ListTenantDomain(ctx request.Ctx)   {}
func (r *R) GetTenantDomain(ctx request.Ctx)    {}
func (r *R) UpdateTenantDomain(ctx request.Ctx) {}
func (r *R) DeleteTenantDomain(ctx request.Ctx) {}

package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
)

func (e *Engine) tryExtractClaim(tenantId string, ctx *gin.Context) (*claim.Session, error) {
	return nil, nil
}

// func (r *runtime) DeriveConsoleClaim(tenantId string, ctx *gin.Context) {
// 	return
// }

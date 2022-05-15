package request

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/temphia/core/backend/server/btypes/models/claim"
)

type Ctx struct {
	Session *claim.Session
	GinCtx  *gin.Context
}

func (c Ctx) Context() context.Context {
	return c.GinCtx.Request.Context()
}

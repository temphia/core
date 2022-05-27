package service

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/core/backend/server/btypes"
)

type TSPRendererBuilder func(app btypes.App) error

type TSPRenderer interface {
	Execute(tenantId, host, file string, ctx *gin.Context)
}

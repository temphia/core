package rtypes

import (
	"github.com/gin-gonic/gin"
)

type Engine interface {
	Run() error
	GetRuntime() Runtime

	OnConsoleExec(tenantId, plugId, agentId, action string, ctx *gin.Context)
	OnRawExec(tenantId, plugId, agentId, action string, ctx *gin.Context)
	OnFedExec(tenantId, resourceId, action string, ctx *gin.Context)
	LaunchSubOrigin(tenantId, plugId, agentId string, ctx *gin.Context)
	LaunchIFrame(tenantId, plugId, agentId string, ctx *gin.Context)

	Serve(tenantId, plugId, agentId, file string, ctx *gin.Context)

	ExecutorFile(tenantId, plugId, agentId, loader string) ([]byte, error)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/backend/server/lib/apiutils"
	"github.com/ztrue/tracerr"
)

func (r *R) WriteJSON(c *gin.Context, resp interface{}, err error) {
	if err != nil {
		tracerr.PrintSourceColor(err)
	}

	apiutils.WriteJSON(c, resp, err)
}

func (r *R) WriteFinal(c *gin.Context, err error) {
	if err != nil {
		tracerr.PrintSourceColor(err)
	}

	apiutils.WriteFinal(c, err)
}

func (r *R) WriteErr(c *gin.Context, msg string) {
	apiutils.WriteErr(c, msg)
}

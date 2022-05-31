package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *R) PortalIndex(ctx *gin.Context) {
	ctx.Writer.Write([]byte(`<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="utf-8" />
	  <meta name="viewport" content="width=device-width,initial-scale=1" />
	  <title>Portal</title>
   <link rel="icon" type="image/png" href="/favicon.png" />
   <link rel="stylesheet"  type="text/css" href="/z/assets/portal.css" />
   <script defer src="/z/assets/portal.js"></script>
   <link rel="stylesheet" href="/z/assets/flatpickr.min.css">
 </head>
 <body></body>
 </html>`))

}

package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *R) AdminRoot2(ctx *gin.Context) {
	ctx.Writer.Write([]byte(`<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="utf-8" />
	  <meta name="viewport" content="width=device-width,initial-scale=1" />
	  <title>Admin Console</title>
   <link rel="icon" type="image/png" href="/favicon.png" />
   <link rel="stylesheet"  type="text/css" href="/assets/console_admin.css" />
   <script defer src="/assets/console_admin.js"></script>
   <link rel="stylesheet" href="/assets/flatpickr.min.css">
 </head>
 <body></body>
 </html>`))

}

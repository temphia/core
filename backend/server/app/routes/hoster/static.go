package hoster

import "github.com/gin-gonic/gin"

func DefaultIndex(ctx *gin.Context) {
	ctx.Writer.Write([]byte(`<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="utf-8" />
	  <meta name="viewport" content="width=device-width,initial-scale=1" />
	  <title>Web Console</title>
	  </script>
   <link rel="icon" type="image/png" href="/favicon.png" />
   <link rel="stylesheet"  type="text/css" href="/z/assets/page_root.css" />
   <script defer src="/z/assets/page_root.js"></script>
 </head>
 <body></body>
 </html>`))
}

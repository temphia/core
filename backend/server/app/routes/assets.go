package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func (r *R) RootIndex(ctx *gin.Context) {
	ctx.Writer.Write([]byte(`<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="utf-8" />
	  <meta name="viewport" content="width=device-width,initial-scale=1" />
	  <title>Web Console</title>
	  </script>
   <link rel="icon" type="image/png" href="/favicon.png" />
   <link rel="stylesheet"  type="text/css" href="./assets/page_root.css" />
   <script defer src="./assets/page_root.js"></script>
 </head>
 <body></body>
 </html>
	  `))
}

func (r *R) NoRoute(c *gin.Context) {
	curPath := c.Request.URL.Path

	if strings.Contains(curPath, "/assets/") {
		pp.Println("@@no_path =>", curPath)

		paths := strings.Split(curPath, "/assets/")
		c.FileFromFS(paths[1], r.assetFS)
		return
	}

	if strings.HasPrefix(curPath, "/console") {
		// fixme => server console root file
		c.Redirect(http.StatusFound, "/console")
	}

}

/*

	/api/:tenant_id/v1
	/public/
	/assets/
	/portal/assets
	/* => root


*/

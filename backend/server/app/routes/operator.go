package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/backend/server/btypes/models/vmodels"
)

func (r *R) OperatorIndex(c *gin.Context) {

	c.Writer.Write([]byte(`<!DOCTYPE html>
	<html lang="en">
	
	<head>
		<meta charset="utf-8" />
		<meta name="viewport" content="width=device-width,initial-scale=1" />
		<title>Operator Console</title>
		<link rel="icon" type="image/png" href="/favicon.png" />
		<link rel="stylesheet" href="./assets/console_operator.css" />
		<script defer src="./assets/console_operator.js"></script>
	</head>
	
	<body>
	
	</body>
	
	</html>`))
}

func (r *R) OperatorAddTenant(c *gin.Context) {
	data := &vmodels.NewTenant{}
	err := c.BindJSON(data)
	if err != nil {
		r.WriteErr(c, err.Error())
		return
	}

	err = r.cOperator.AddTenant(data)
	r.WriteFinal(c, err)
}

func (r *R) OperatorUpdateTenant(c *gin.Context) {
	data := make(map[string]interface{})
	err := c.BindJSON(&data)
	if err != nil {
		r.WriteErr(c, err.Error())
		return
	}

	slug := data["slug"].(string)
	delete(data, "slug")

	err = r.cOperator.UpdateTenant(slug, data)
	r.WriteFinal(c, err)

}

func (r *R) OperatorListTenant(c *gin.Context) {
	data, err := r.cOperator.ListTenant()
	r.WriteJSON(c, data, err)
}

type DelTenant struct {
	Slug string `json:"slug,omitempty"`
}

func (r *R) OperatorDeleteTenant(c *gin.Context) {
	dt := &DelTenant{}

	err := c.BindJSON(&dt)
	if err != nil {
		r.WriteErr(c, err.Error())
		return
	}
	if dt.Slug == "" {
		r.WriteErr(c, "empty tenant slug")
		return
	}

	r.cOperator.DeleteTenant(dt.Slug)

}

func (r *R) OperatorStats(c *gin.Context) {

}

func (r *R) OperatorTenantToken(c *gin.Context) {

}

func (r *R) OperatorLogin(c *gin.Context) {
	data := &vmodels.OperatorLoginReq{}

	err := c.BindJSON(data)
	if err != nil {
		r.WriteErr(c, err.Error())
		return
	}

	resp, err := r.cOperator.Login(data)
	r.WriteJSON(c, resp, err)
}

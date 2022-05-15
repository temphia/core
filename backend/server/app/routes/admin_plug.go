package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/backend/server/btypes/models/entities"
	"github.com/temphia/temphia/backend/server/lib/apiutils/request"
)

func (r *R) NewPlug(ctx request.Ctx) {
	data := &entities.Plug{}

	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	data.TenantId = ctx.Session.TenentID

	err = r.cAdmin.PlugNew(ctx.Session, data)
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) UpdatePlug(ctx request.Ctx) {
	data := make(map[string]interface{})
	err := ctx.GinCtx.BindJSON(&data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.cAdmin.PlugUpdate(ctx.Session, ctx.GinCtx.Param("plug_id"), data)

}

func (r *R) GetPlug(ctx request.Ctx) {
	plug, err := r.cAdmin.PlugGet(ctx.Session, ctx.GinCtx.Param("plug_id"))
	r.WriteJSON(ctx.GinCtx, plug, err)
}
func (r *R) DelPlug(ctx request.Ctx) {
	err := r.cAdmin.PlugDel(ctx.Session, ctx.GinCtx.Param("plug_id"))
	r.WriteJSON(ctx.GinCtx, nil, err)
}

func (r *R) ListPlug(ctx request.Ctx) {
	pgs, err := r.cAdmin.PlugList(ctx.Session)
	r.WriteJSON(ctx.GinCtx, pgs, err)
}

func (r *R) NewAgent(ctx request.Ctx) {
	data := &entities.Agent{}
	err := ctx.GinCtx.BindJSON(data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	data.TenantId = ctx.Session.TenentID

	err = r.cAdmin.AgentNew(ctx.Session, data)
	r.WriteFinal(ctx.GinCtx, err)
}

func (r *R) UpdateAgent(ctx request.Ctx) {
	data := make(map[string]interface{})
	err := ctx.GinCtx.BindJSON(&data)
	if err != nil {
		r.WriteErr(ctx.GinCtx, err.Error())
		return
	}

	r.cAdmin.AgentUpdate(ctx.Session, ctx.GinCtx.Param("plug_id"), ctx.GinCtx.Param("agent_id"), data)
}

func (r *R) GetAgent(ctx request.Ctx) {
	agent, err := r.cAdmin.AgentGet(ctx.Session, ctx.GinCtx.Param("plug_id"), ctx.GinCtx.Param("agent_id"))
	r.WriteJSON(ctx.GinCtx, agent, err)
}

func (r *R) DelAgent(ctx request.Ctx) {
	err := r.cAdmin.AgentDel(ctx.Session, ctx.GinCtx.Param("plug_id"), ctx.GinCtx.Param("agent_id"))
	r.WriteFinal(ctx.GinCtx, err)
}

func (r *R) ListAgent(ctx request.Ctx) {
	agents, err := r.cAdmin.AgentList(ctx.Session, ctx.GinCtx.Param("plug_id"))
	r.WriteJSON(ctx.GinCtx, agents, err)
}

func (r *R) PairAgentToken(ctx request.Ctx) {

}

var DefaultPlugIcon = []byte(`<svg
xmlns="http://www.w3.org/2000/svg"
class="w-28 h-28 text-gray-500"
fill="none"
viewBox="0 0 24 24"
stroke="currentColor"
>
<path
  stroke-linecap="round"
  stroke-linejoin="round"
  stroke-width="2"
  d="M11 4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a2 2 0 110-4h1a1 1 0 001-1V7a1 1 0 011-1h3a1 1 0 001-1V4z"
/>
</svg>`)

func (p *R) PlugIcon(c *gin.Context) {
	// c.Param("plug_id")
	c.Writer.Write(DefaultPlugIcon)
}

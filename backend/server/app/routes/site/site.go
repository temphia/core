package site

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/store"
)

type Manager struct {
	cabhub     store.CabinetHub
	corehub    store.CoreHub
	revDomains map[string]string
}

func NewManager(cabhub store.CabinetHub, corehub store.CoreHub) Manager {
	return Manager{
		cabhub:     nil,
		corehub:    nil,
		revDomains: make(map[string]string),
	}

}

func (m *Manager) ServeIndex(c *gin.Context) {
	m.serve("index.html", c)
}

func (m *Manager) ServeAny(file string, c *gin.Context) {
	m.serve(file, c)
}

func (m *Manager) serve(file string, c *gin.Context) {
	tenantId, _ := c.Cookie("tenant_id")
	if tenantId == "" {
		tenantId = c.Query("tenant_id")
	}

	switch c.Request.URL.Hostname() {
	case "localhost", "127.0.0.1":
		if tenantId == "" {
			DefaultIndex(c)
			return
		}
		m.writeFile(c, tenantId, "", "public", file)
	default:
		if tenantId == "" {
			tenantId = m.lookupTenant(c.Request.URL.Hostname())
			if tenantId == "" {
				if file == "index.html" {
					DefaultIndex(c)
					return
				}
				return
			}
		}
		m.writeFile(c, tenantId, "", "public", file)
	}

}

func (m *Manager) writeFile(c *gin.Context, tenantId, source, folder, file string) {

	var csource store.CabinetSourced

	if source == "" {
		csource = m.cabhub.Default(tenantId)
	} else {
		csource = m.cabhub.GetSource(tenantId, source)
	}

	out, err := csource.GetBlob(c, folder, file)
	if err != nil {
		pp.Println(err)
		return
	}
	c.Writer.Write(out)
}

func (m *Manager) lookupTenant(host string) string {
	tenant, ok := m.revDomains[host]
	if ok {
		return tenant
	}

	// fixme => dns txt loopup here [ temphia_clusterxxx = <tokenxyz> ]

	return ""
}

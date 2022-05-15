package local

import (
	"github.com/temphia/temphia/backend/server/registry"
	"github.com/temphia/temphia/data"
)

func init() {
	registry.SetRepoBuilder("embed", data.NewEmbed)
}

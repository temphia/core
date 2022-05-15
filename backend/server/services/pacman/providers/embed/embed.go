package local

import (
	"github.com/temphia/core/backend/server/registry"
	"github.com/temphia/core/data"
)

func init() {
	registry.SetRepoBuilder("embed", data.NewEmbed)
}

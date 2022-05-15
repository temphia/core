package data

import (
	"embed"
)

//go:embed schema
//go:embed templates
//go:embed assets
//go:embed repo
var dataDir embed.FS

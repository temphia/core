package renderer

import (
	"io/fs"

	"github.com/flosch/pongo2/v5"
	"github.com/temphia/core/backend/server/btypes/service"
)

type RenderOptions struct {
	Globals       map[string]interface{}
	Theme         string
	OverrideTheme string
}

var _ fs.FS = (*DomainRenderer)(nil)

type DomainRenderer struct {
	pacman              service.Pacman
	Template            *pongo2.TemplateSet
	ThemeBprint         string
	OverrideThemeBprint string

	// wizards             map[string]int64
}

func (d *DomainRenderer) Open(name string) (fs.File, error) {
	return &File{
		name:   name,
		parent: d,
	}, nil
}

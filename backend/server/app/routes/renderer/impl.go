package renderer

import (
	"github.com/flosch/pongo2/v5"
	"github.com/k0kubun/pp"
)

func (r *Renderer) domainRenderer(domain string) (*DomainRenderer, error) {
	r.mlock.Lock()
	defer r.mlock.Unlock()

	dr, ok := r.renderers[domain]
	if ok {
		return dr, nil
	}

	opts, err := r.domainRenderOptions(domain)
	if err != nil {
		return nil, err
	}

	dr = &DomainRenderer{
		pacman:              r.pacman,
		Template:            nil,
		ThemeBprint:         opts.Theme,
		OverrideThemeBprint: opts.OverrideTheme,
		RootBprint:          "",
		RootPrefix:          "",
	}

	loader := pongo2.NewFSLoader(dr)
	tplset := pongo2.NewSet(domain, loader)
	tplset.Globals.Update(opts.Globals)

	return dr, nil

}

func (r *Renderer) domainRenderOptions(domain string) (*RenderOptions, error) {
	return nil, nil
}

func (r *Renderer) render(domain string) ([]byte, error) {
	dr, err := r.domainRenderer(domain)
	if err != nil {
		return nil, err
	}

	pp.Println(dr)

	dr.Template.RenderTemplateFile("file", nil)

	return nil, nil
}

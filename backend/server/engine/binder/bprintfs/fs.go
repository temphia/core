package bprintfs

import (
	"io/fs"

	"github.com/temphia/core/backend/server/btypes/rtypes"
)

var _ fs.FS = (*FS)(nil)

type FS struct {
	b rtypes.Bindings

	// files []string
}

func New(b rtypes.Bindings) *FS {
	return &FS{b: b}
}

func (s *FS) Open(name string) (fs.File, error) {
	return &File{
		name: name,
		b:    s.b,
	}, nil
}

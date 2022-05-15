package binder

import (
	"bytes"
	"html/template"

	"github.com/temphia/temphia/backend/server/btypes/rtypes"
)

func (b *Binder) SelfGetFile(file string) ([]byte, error) {
	return b.factory.Pacman.BprintGetBlob(b.namespace, b.job.Plug.BprintId, file)
}

func (b *Binder) SelfAddFile(file string, data []byte) error {
	return b.factory.Pacman.BprintNewBlob(b.namespace, b.job.Plug.BprintId, file, data)
}
func (b *Binder) SelfUpdateFile(file string, data []byte) error {
	return b.factory.Pacman.BprintUpdateBlob(b.namespace, b.job.Plug.BprintId, file, data)
}

func (b *Binder) SelfListResources() ([]*rtypes.Resource, error) {
	return nil, nil
}
func (b *Binder) SelfGetResource(name string) (*rtypes.Resource, error) {
	return nil, nil
}

func (b *Binder) SelfIncomingConns() ([]*rtypes.Connection, error) {
	return nil, nil
}

func (b *Binder) SelfOutgoingConns() ([]*rtypes.Connection, error) {
	return nil, nil
}

func (b *Binder) SelfExecuteSlot(name string, opts rtypes.SlotOption) ([]byte, error) {
	if opts.Async {
		go b.execslot(name, opts)
		return nil, nil
	}

	return b.execslot(name, opts)
}

func (b *Binder) SelfRenderFile(file string, opts rtypes.RenderOption) ([]byte, error) {
	_, err := b.GetSelfFile(file)
	if err != nil {
		return nil, err
	}

	adapter := SelfFs{
		b: b,
	}

	tpl, err := template.New("plug ").ParseFS(&adapter, file)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, opts.Data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (b *Binder) execslot(name string, opts rtypes.SlotOption) ([]byte, error) {
	return nil, nil

}

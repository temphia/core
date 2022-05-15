package binder

import (
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
)

func (b *Binder) GetPlugKVBindings() rtypes.PlugKVBindings {
	return &b.plugKvBinder
}

func (b *Binder) GetCabinetBindings() rtypes.CabinetBindings {
	return b
}

func (b *Binder) GetSockdBindings() rtypes.SockdBindings { return b }
func (b *Binder) GetUserBindings() rtypes.UserBindings   { return b }
func (b *Binder) GetSelfBindings() rtypes.SelfBindings   { return b }
func (b *Binder) GetPCache() rtypes.PCache               { return b }

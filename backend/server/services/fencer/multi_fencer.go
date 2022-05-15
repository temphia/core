package fencer

import (
	"github.com/temphia/temphia/backend/server/btypes"
	"github.com/temphia/temphia/backend/server/btypes/env"
	"github.com/temphia/temphia/backend/server/btypes/service"
)

var _ service.Fencer = (*MultiFencer)(nil)

type MultiFencer struct {
	tenantFencers map[string]*fencer
}

func New(_app btypes.App) *MultiFencer {
	return &MultiFencer{
		tenantFencers: make(map[string]*fencer),
	}
}

func (f *MultiFencer) Name() string { return "fencer" }
func (f *MultiFencer) Init() error  { return nil }
func (f *MultiFencer) Start() error { return nil }
func (f *MultiFencer) Stop() error  { return nil }

func (f *MultiFencer) ResourcePolicy(string, *env.ResourceEval) error {
	return nil
}

func (f *MultiFencer) SignalPolicy(string, *env.Signal) error {
	return nil
}

func (f *MultiFencer) AdviseryPolicy(string, *env.Advisery) error {
	return nil
}

func (f *MultiFencer) AuthFlowCheck(string, *env.AuthFlow) error {
	return nil
}

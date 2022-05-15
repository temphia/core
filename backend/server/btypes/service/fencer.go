package service

import "github.com/temphia/temphia/backend/server/btypes/env"

type Fencer interface {
	Service
	ExecutorFencer
	AuthFencer
}

type ExecutorFencer interface {
	AdviseryFencer
	ResourcePolicy(string, *env.ResourceEval) error
	SignalPolicy(string, *env.Signal) error
}

type AdviseryFencer interface {
	AdviseryPolicy(string, *env.Advisery) error
}

type AuthFencer interface {
	AuthFlowCheck(string, *env.AuthFlow) error
}

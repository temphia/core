package fencer

import (
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

// policy is wrapper around vm.Program with premature
// optimcation of `true` or `false` evaluates directly without evaluating
// actaual VM
type policy struct {
	inner       *vm.Program
	skipped     bool
	staticValue bool
}

func newPolicy(code string, ctxEnv interface{}) policy {
	if code == `true` {
		return policy{
			inner:       nil,
			skipped:     true,
			staticValue: true,
		}

	} else if code == `false` {
		return policy{
			inner:       nil,
			skipped:     true,
			staticValue: false,
		}
	}

	// ctx given at build time is just for typechecking variables
	// actual ctx has to given when executing again
	program, err := expr.Compile(code, expr.Env(ctxEnv))
	if err != nil {
		panic(err)
	}

	return policy{
		inner:       program,
		skipped:     false,
		staticValue: false,
	}
}

func (p policy) Execute(env interface{}) (interface{}, error) {
	if p.skipped {
		return p.staticValue, nil
	}
	return expr.Run(p.inner, env)
}

func (p policy) isExpty() bool {
	return p.inner == nil
}

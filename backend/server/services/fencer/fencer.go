package fencer

type fencer struct {
	parent        *MultiFencer
	tenantId      string
	groupPolicies map[string]*userGroupPolicy
	globalPolicy  pool
	agentsPool    map[string]pool
}

func (f *fencer) Execute(ctx interface{}) (interface{}, error) {
	program := f.globalPolicy.Borrow()
	if program.isExpty() {
		program = f.newPolicy(ctx)
	}

	return program.Execute(ctx)
}

func (f *fencer) newPolicy(ctxEnv interface{}) policy {
	code := `true`
	return newPolicy(code, ctxEnv)
}

// userGroupPolicy
type userGroupPolicy struct {
	authPolicy   pool
	loginPolicy  pool
	signUpPolicy pool
	blobPolicy   pool
}

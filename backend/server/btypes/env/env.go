package env

import (
	claim "github.com/temphia/temphia/backend/server/btypes/models/claim"
	"gitlab.com/mr_balloon/golib/hmap"
)

type Blob struct {
	ContextVars  map[string]interface{}
	SessionClaim *claim.Session
}

type Advisery struct {
	ExecEnv     interface{}
	Inner       hmap.H
	ContextVars map[string]string
}

package env

import (
	"github.com/temphia/core/backend/server/btypes/models/claim"
	"github.com/temphia/core/backend/server/btypes/models/entities"
)

type (
	AuthFlow struct {
		Stage         string
		User          *entities.User
		UserGroup     *entities.UserGroup
		UserClaim     *claim.User
		SiteClaim     *claim.Site
		ChallengeType string
		ChallengeData string
		FlowSideEffects
	}

	FlowSideEffects struct {
		ErrRedirrect  string // to url
		Message       string
		ExtraAttrs    map[string]string
		CopyAttrs     []string
		SkipChallenge bool
		Err           bool
	}
)

func (af *AuthFlow) ErrRedirrect(url string) {
	af.FlowSideEffects.ErrRedirrect = url
	af.Err = true
}

func (af *AuthFlow) SetMessage(msg string) {
	af.FlowSideEffects.Message = msg
}

func (af *AuthFlow) SetErr(reason string) {
	af.FlowSideEffects.Message = reason
	af.Err = true
}

func (af *AuthFlow) SetAttr(key, value string) {
	af.FlowSideEffects.ExtraAttrs[key] = value
}

func (af *AuthFlow) CopyAttrFromToken(key string) {
	af.CopyAttrs = append(af.CopyAttrs, key)
}

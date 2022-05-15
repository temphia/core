package env

import (
	"github.com/temphia/temphia/backend/server/btypes/models/claim"
)

type Service struct {
	UserClaim *claim.User
	Type      string
	Source    string
	Object    string
	Err       string
	// side effects
	CopyAttrs []string
}

func (s *Service) Ok() bool {
	return s.Err == ""
}

func (s *Service) SetAttrCopy(key string) {
	s.CopyAttrs = append(s.CopyAttrs, key)
}

func (s *Service) Raise(msg string) {
	s.Err = msg
}

func (s *Service) SetAttr(key, val string) {
	s.CopyAttrs = append(s.CopyAttrs, key)
	s.UserClaim.Attributes[key] = val
}

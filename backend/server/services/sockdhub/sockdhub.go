package sockdhub

import (
	"github.com/temphia/temphia/backend/server/btypes/service"
)

type SockdHub struct {
	sockd service.SockCore
}

func New(sockd service.SockCore) *SockdHub {
	return &SockdHub{
		sockd: sockd,
	}
}

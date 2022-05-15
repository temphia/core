package sockcore

import (
	"encoding/json"
)

const (
	// fixme => system message

	MESSAGE_SERVER_DIRECT    = "server_direct"
	MESSAGE_SERVER_BROADCAST = "server_broadcast"
	MESSAGE_SERVER_PUBLISH   = "server_publish"
	MESSAGE_PEER_DIRECT      = "peer_direct"
	MESSAGE_PEER_BROADCAST   = "peer_broadcast"
	MESSAGE_PEER_PUBLISH     = "peer_publish"
)

type message struct {
	Room        string          `json:"room,omitempty"`
	FromId      string          `json:"from_id,omitempty"`
	Type        string          `json:"type,omitempty"`
	ServerIdent string          `json:"server_ident,omitempty"`
	XID         string          `json:"xid,omitempty"`
	Ticket      string          `json:"ticket,omitempty"`
	Targets     []string        `json:"targets,omitempty"`
	Payload     json.RawMessage `json:"payload,omitempty"`
}

func (m *message) JSON() ([]byte, error) {
	return json.Marshal(m)
}

type TicketParse func(ns, room, connId, ticket string, rtargets []string) ([]string, error)

type ticketClaim struct {
	partial    bool
	targetType string
	targets    []string
}

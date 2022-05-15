package service

type Signer interface {
	GlobalSignRaw(payload string) (string, error)
	GlobalParseRaw(token string) (string, error)
	Sign(ns string, payload interface{}) (string, error)
	Parse(ns string, token string, target interface{}) error
	SignRaw(ns string, payload string) (string, error)
	ParseRaw(ns string, token string) (string, error)
}

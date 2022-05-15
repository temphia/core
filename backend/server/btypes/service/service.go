package service

type Service interface {
	Component
	Start() error
	Stop() error
}

package service

// component is a component that takes dependencies dynamically and
type Component interface {
	Name() string
	Init() error
}

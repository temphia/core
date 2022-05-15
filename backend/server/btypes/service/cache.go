package service

type Cache interface {
	Put(key, value string) error
	Get(key string) (string, error)
	Expire(key string) error
}

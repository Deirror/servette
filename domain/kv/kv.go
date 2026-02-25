package kv

type Storer interface {
	Set(key string, value string, ttlSeconds int) error
	Get(key string) (string, error)
	Delete(key string) error
	Exists(key string) (int64, error)
}

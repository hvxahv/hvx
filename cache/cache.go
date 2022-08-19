package cache

type Cache interface {
	// String operations.
	EXISTS(key string) (bool, error)
	Get(key string) (string, error)
	Set(key string, value string) error
	Del(key string) error

	// Hash operations.
	HEXISTS(key string, field string) (bool, error)
	HSET(key string, field string, value string) error
	HGET(key string, field string) (string, error)
	HGETALL(key string) (map[string]string, error)
	HMSET(key string, fields map[string]string) error
	HMGET(key string, fields []string) ([]string, error)
}

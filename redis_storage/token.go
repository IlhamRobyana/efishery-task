package redis_storage

import (
	"github.com/gomodule/redigo/redis"
)

type TokenStorage struct {
}

func (t *TokenStorage) IsAvailable(token string) bool {
	conn := getPool().Get()
	defer conn.Close()

	if _, err := redis.String(conn.Do("GET", token)); err == nil {
		return true
	}
	return false
}

func (t *TokenStorage) Create(token string) (err error) {
	conn := getPool().Get()
	defer conn.Close()

	_, err = redis.String(conn.Do("SET", token, true))
	return
}

func (t *TokenStorage) Update(token string, oldToken string) error {
	return nil
}

func (t *TokenStorage) Delete(token string) error {
	return nil
}

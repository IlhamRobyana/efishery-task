package redis_storage

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

var pool *redis.Pool

func getPool() (redisPool *redis.Pool) {

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("can't load .env : %v", err))
	}

	redisAddress := os.Getenv("REDIS_ADDRESS")

	if pool == nil {
		pool = &redis.Pool{
			MaxIdle:   80,
			MaxActive: 12000,
			Dial: func() (conn redis.Conn, err error) {
				conn, err = redis.Dial("tcp", redisAddress)
				return
			},
		}
	}
	redisPool = pool
	return
}

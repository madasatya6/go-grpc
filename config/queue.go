package config

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"

	"go_grpc/lib"
)

func NewQueue() *lib.Queue {
	redisPool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")))
		},
	}

	return lib.NewQueue(redisPool)
}

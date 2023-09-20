package config

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go_grpc/lib"
	"os"
)

func NewPubSub() *lib.PubSub {
	redisConn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")))
	if err != nil {
		panic(err)
	}

	return &lib.PubSub{Redis: redisConn}
}

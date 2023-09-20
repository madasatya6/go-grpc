package lib

import "github.com/gomodule/redigo/redis"

type PubSub struct {
	Redis redis.Conn
}

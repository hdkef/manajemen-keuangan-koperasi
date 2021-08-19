package driver

import (
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

// var REDISPASS string
var REDISADDR string

type RedisDriver struct {
	C redis.Conn
}

func init() {
	_ = godotenv.Load()
	REDISADDR = os.Getenv("REDISADDR")
	// REDISPASS = os.Getenv("REDISPASS")
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDISADDR)
			if err != nil {
				panic(err.Error())
			}
			// _, err = c.Do("AUTH", REDISPASS)
			// if err != nil {
			// 	panic(err.Error())
			// }
			return c, err
		},
	}
}

func RedisConn() *RedisDriver {

	return &RedisDriver{
		C: newRedisPool().Get(),
	}
}

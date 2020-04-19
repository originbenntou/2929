package client

import (
	gredis "github.com/go-redis/redis"
)

var Client *gredis.Client

func init() {
	Client = gredis.NewClient(&gredis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

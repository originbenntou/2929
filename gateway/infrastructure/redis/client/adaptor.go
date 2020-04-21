package client

import (
	gredis "github.com/go-redis/redis/v7"
	"github.com/originbenntou/2929BE/shared/logger"
)

const EMPTY = gredis.Nil

var Client *gredis.Client

func init() {
	Client = gredis.NewClient(&gredis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := Client.Ping().Result()
	if err != nil {
		logger.Common.Error(err.Error())
	}

}

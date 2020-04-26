package client

import (
	gredis "github.com/go-redis/redis/v7"
	"github.com/originbenntou/2929BE/shared/logger"
)

const EMPTY = gredis.Nil

// ログインチェック
// Hash型 key: Token, field: uid,cid
var TokenClient *gredis.Client

// FIXME: configをつかって書き直す
func init() {
	TokenClient = gredis.NewClient(&gredis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Redis疎通確認
	var err error

	_, err = TokenClient.Ping().Result()
	if err != nil {
		logger.Common.Error(err.Error())
	}
}

package client

import (
	gredis "github.com/go-redis/redis/v7"
	"github.com/originbenntou/2929BE/shared/logger"
)

const EMPTY = gredis.Nil

// ログインチェック
// Hash型 key: Token, field: uid,cid
var TokenClient *gredis.Client

// uid重複チェック
// List型 key: user,  list: uid
var UidClient *gredis.Client

// cid重複チェック
// List型 key: company, list: cid
var CidClient *gredis.Client

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

	UidClient = gredis.NewClient(&gredis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	_, err = UidClient.Ping().Result()
	if err != nil {
		logger.Common.Error(err.Error())
	}

	CidClient = gredis.NewClient(&gredis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       2,  // use default DB
	})

	_, err = CidClient.Ping().Result()
	if err != nil {
		logger.Common.Error(err.Error())
	}
}

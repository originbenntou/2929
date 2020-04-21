package middleware

import (
	"fmt"
	"github.com/google/uuid"
	redis "github.com/originbenntou/2929BE/gateway/infrastructure/redis/client"
	"github.com/originbenntou/2929BE/gateway/interfaces/support"
	"github.com/originbenntou/2929BE/shared/logger"
	"go.uber.org/zap"
	"net/http"
)

const (
	XRequestIDKey = "X-Request-Id"
	InValidCookie = `{
  "error": {
    "errors": [
      {
        "message": "Invalid cookie"
      }
    ],
    "data": null
  }
}`
)

func Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(XRequestIDKey)
		if traceID == "" {
			traceID = newTraceID()
		}
		ctx := support.AddTraceIDToContext(r.Context(), traceID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func newTraceID() string {
	return uuid.New().String()
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			logger.Common.Info(
				"OK",
				zap.String("TraceID", support.GetTraceIDFromContext(r.Context())),
				zap.String("Method", r.Method),
				zap.String("Request", r.RequestURI),
			)
		}()
		next.ServeHTTP(w, r)
	})
}

func NewAuthentication() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uid, err := redis.Client.HGet("6b8a733-9d75-44c3-9cc9-34addc8f965a", "uid").Result()
			if uid == "" || err == redis.EMPTY {
				logger.Common.Info(err.Error())
				http.Error(w, InValidCookie, http.StatusForbidden)
				return
			}

			cid, err := redis.Client.HGet("6b85a733-9d75-44c3-9cc9-34addc8f965a", "cid").Result()
			if cid == "" || err == redis.EMPTY {
				logger.Common.Info(err.Error())
				http.Error(w, InValidCookie, http.StatusForbidden)
				return
			}

			fmt.Println(uid, cid, err)

			//ctx := support.AddUserToContext(r.Context(), resp.User)
			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}

package middleware

import (
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
			defer func() {
				//if err := redis.Client.Close(); err != nil {
				//	logger.Common.Info(err.Error())
				//	http.Error(w, InValidCookie, http.StatusForbidden)
				//	return
				//}
			}()

			// FIXME: tokenが存在すればOKにする
			uid, err := redis.TokenClient.HGet("d0b50ce2-49f4-4cb7-b0e6-6acfc1d98a0a", "uid").Result()
			if uid == "" || err == redis.EMPTY {
				logger.Common.Info(err.Error())
				http.Error(w, InValidCookie, http.StatusForbidden)
				return
			}

			cid, err := redis.TokenClient.HGet("d0b50ce2-49f4-4cb7-b0e6-6acfc1d98a0a", "cid").Result()
			if cid == "" || err == redis.EMPTY {
				logger.Common.Info(err.Error())
				http.Error(w, InValidCookie, http.StatusForbidden)
				return
			}

			//ctx := support.AddUserToContext(r.Context(), resp.User)
			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}

func NewCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// FIXME: フロントサーバー（とマネージャーサーバー）のオリジンで制限を掛ける
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}

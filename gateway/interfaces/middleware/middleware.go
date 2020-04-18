package middleware

import (
	"github.com/google/uuid"
	"github.com/originbenntou/2929BE/gateway/interfaces/support"
	"net/http"
)

const (
	xRequestIDKey = "X-Request-Id"
)

func Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(xRequestIDKey)
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

//func NewAuthentication(
//	userClient pbUser.UserServiceClient,
//	sessionStore session.Store) func(next http.HandlerFunc) http.HandlerFunc {
//	return func(next http.HandlerFunc) http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			sessionID := session.GetSessionIDFromRequest(r)
//			v, ok := sessionStore.Get(sessionID)
//			if !ok {
//				http.Redirect(w, r, "/login", http.StatusFound)
//				return
//			}
//			userID, ok := v.(uint64)
//			if !ok {
//				http.Redirect(w, r, "/login", http.StatusFound)
//				return
//			}
//			ctx := r.Context()
//			resp, err := userClient.FindUser(ctx, &pbUser.FindUserRequest{
//				UserId: userID,
//			})
//			if err != nil {
//				http.Redirect(w, r, "/login", http.StatusFound)
//				return
//			}
//			ctx = support.AddUserToContext(ctx, resp.User)
//			next.ServeHTTP(w, r.WithContext(ctx))
//		}
//	}
//}

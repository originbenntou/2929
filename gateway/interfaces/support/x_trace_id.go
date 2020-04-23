package support

import (
	"context"
)

type contextKeyTraceID struct{}

func GetTraceIDFromContext(ctx context.Context) string {
	id := ctx.Value(contextKeyTraceID{})
	traceID, ok := id.(string)
	if !ok {
		return ""
	}
	return traceID
}

func AddTraceIDToContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, contextKeyTraceID{}, traceID)
}

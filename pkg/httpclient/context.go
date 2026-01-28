package httpclient

import (
	"context"

	"github.com/gin-gonic/gin"
)

// ContextWithRequestID creates a context with the request ID from gin context
// Use this when you need to make outbound HTTP calls from a handler
func ContextWithRequestID(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	if requestID := c.GetString(RequestIDKey); requestID != "" {
		ctx = context.WithValue(ctx, RequestIDKey, requestID)
	}
	return ctx
}

// GetRequestIDFromContext retrieves request ID from context
func GetRequestIDFromContext(ctx context.Context) string {
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		return requestID.(string)
	}
	return ""
}

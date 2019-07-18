package golang_sdk

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinRequestContext gin.Context

func (c GinRequestContext) GetRequest() *http.Request {
	return c.Request
}

func (c GinRequestContext) GetWriter() http.ResponseWriter {
	return c.Writer
}

func FromGinContext(ctx *gin.Context) *ApiContext {
	c := (*GinRequestContext)(ctx)
	return &ApiContext{
		requestContext: c,
	}
}

type GinHandlerFunc func(ctx *gin.Context, apiCtx *ApiContext)

func Gin(fun GinHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiCtx := FromGinContext(ctx)
		fun(ctx, apiCtx)
	}
}

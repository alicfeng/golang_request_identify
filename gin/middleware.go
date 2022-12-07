package gin_request_identify

import (
	request "github.com/alicfeng/golang_request_identify"
	"github.com/gin-gonic/gin"
)

// Server 追踪请求链路服务端中间件.
func Server(context *gin.Context) {
	if "" == context.Request.Header.Get(request.HeaderFieldRequestIdentity) {
		context.Request.Header.Set(request.HeaderFieldRequestIdentity, request.Create())
	}

	context.Next()
}

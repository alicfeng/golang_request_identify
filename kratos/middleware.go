package kratos_request_identify

import (
	"context"
	request "github.com/alicfeng/golang_request_identify"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// Server 追踪请求链路服务端中间件.
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
			// 1.服务启用metadata且请求无携带追踪标识 则生成唯一标识并应用
			if md, ok := metadata.FromServerContext(ctx); ok && "" == md.Get(request.HeaderFieldRequestIdentity) {
				md.Set(request.HeaderFieldRequestIdentity, request.Create())
			}

			// 2.设置http响应头 仅对http通信生效 告知客户端追踪标识的编码
			if tr, ok := transport.FromServerContext(ctx); ok && transport.KindHTTP == tr.Kind() {
				if ht, success := tr.(*http.Transport); success {
					ht.ReplyHeader().Set(request.HeaderFieldRequestIdentity, request.Value(ctx))
				}
			}

			return handler(ctx, req)
		}
	}
}

// Client 追踪请求链路客户端中间件.
func Client() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {

			// 1.将追踪标识追加到上下文中
			ctx = metadata.AppendToClientContext(ctx, request.HeaderFieldRequestIdentity, request.Value(ctx))

			return handler(ctx, req)
		}
	}
}

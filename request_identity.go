package request_identify

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/gofrs/uuid"
	"strings"
)

const (
	HeaderFieldRequestIdentity = "x-md-global-request-identity" // 请求头追踪标识字段名称
)

// Create 生成请求唯一标识.
func Create() string {
	return strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(uuid.Must(uuid.NewV4()).String()))))
}

// Value 获取请求标识.
func Value(ctx context.Context) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(HeaderFieldRequestIdentity)
	}

	return ""
}

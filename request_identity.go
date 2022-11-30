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

// CreateRequestIdentity 生成请求唯一标识.
func CreateRequestIdentity() string {
	return strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(uuid.Must(uuid.NewV4()).String()))))
}

// GetRequestIdentity 获取请求标识.
func GetRequestIdentity(ctx context.Context) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(HeaderFieldRequestIdentity)
	}

	return ""
}

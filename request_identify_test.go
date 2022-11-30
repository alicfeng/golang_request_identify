package request_identify

import (
	"testing"
)

func TestCreateRequestIdentity(t *testing.T) {
	cv := CreateRequestIdentity()

	// 1.生成标识编码不允许为空
	if "" == cv {
		t.Error("expected non null value, but result is empty")
	}

	// 2.生成标识编码长度为32
	if 32 != len(cv) {
		t.Errorf("expected length 32,but result is %d", len(cv))
	}
}

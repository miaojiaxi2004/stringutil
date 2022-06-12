package stringutil

import (
	"testing"
)

// 测试函数翻转功能是否有效
func TestReverse(t *testing.T) {
	name := "Test"
	msg, err := Reverse("tseT")
	if err != nil || name != msg {
		t.Fatalf(`Reverse("") = %q, %v, want "", error`, msg, err)
	}
}

// 测试我们的代码是否正确通过字符串为空得到报错结果
func TestReverseEmpty(t *testing.T) {
	msg, err := Reverse("")
	if msg != "" || err == nil {
		t.Fatalf(`Reverse("") = %q, %v, want "", error`, msg, err)
	}
}

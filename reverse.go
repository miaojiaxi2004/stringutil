package stringutil

import "errors"

// 将字符串翻转
func Reverse(s string) (string, error) {
	if s == "" {
		return s, errors.New("请输入字符串")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

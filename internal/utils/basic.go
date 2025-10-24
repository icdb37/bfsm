package utils

import "strings"

// PstrTrims 去除空白字符
func PstrTrims(strs ...*string) {
	for _, str := range strs {
		if str != nil {
			*str = strings.TrimSpace(*str)
		}
	}
}

// SstrTrims 字符串处理
func SstrTrims(strs []string) {
	for i, size := 0, len(strs); i < size; i++ {
		strs[i] = strings.TrimSpace(strs[i])
	}
}

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

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Hash 计算哈希
func Hash(ss ...string) string {
	md5 := md5.New()
	for _, s := range ss {
		md5.Write([]byte(s))
	}
	return hex.EncodeToString(md5.Sum(nil))
}

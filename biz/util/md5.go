package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var salt = "hertz_basic_huahuo"

// HashStringWithSalt 对输入字符串和盐值进行 MD5 加密，并返回十六进制字符串
func MD5Hash(s string) string {
	// 将输入字符串和盐值连接起来
	saltedInput := fmt.Sprintf("%s%s", s, salt)

	// 创建一个新的 MD5 哈希实例
	h := md5.New()

	// 将加盐后的字符串写入哈希实例
	h.Write([]byte(saltedInput))

	// 计算 MD5 校验和
	checksum := h.Sum(nil)

	// 将校验和转换为十六进制字符串
	return hex.EncodeToString(checksum)
}

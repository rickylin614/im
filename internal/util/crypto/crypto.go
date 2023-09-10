package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

const passKey = "passsss"

func Hash(str string) string {
	// 创建一个 MD5 哈希对象
	hasher := md5.New()

	// 将字符串转换为字节数组并写入哈希对象
	hasher.Write([]byte(str + passKey))

	// 计算 MD5 哈希值
	hashBytes := hasher.Sum(nil)

	// 将哈希值转换为十六进制字符串
	md5String := hex.EncodeToString(hashBytes)
	return md5String
}

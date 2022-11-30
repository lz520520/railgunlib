package goutils

import (
	"math/rand"
	"time"
)

const (
	hexMeta = "0123456789abcdef"

	AsciiLitter = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	UpperLitter = "QWERTYUIOPASDFGHJKLZXCVBNM"
	LowerLitter = "qwertyuiopasdfghjklzxcvbnm"
	Digits      = "1234567890"
)

// 随机生成 MD5 HASH 值
func RandomMD5Hash() string {
	return randStrWithMeta(32, hexMeta)
}

// 随机生成指定长度的字符串
func RandomHexString(size int) (ret string) {
	return randStrWithMeta(size, hexMeta)
}

// 随机int
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

// 随机数字字符串
func RandDigital(n int) string {
	if n == 0 {
		n = RandInt(3, 20)
	}
	return RandStrWithMeta(n, Digits)
}

// 自定义meta的随机
func RandStrWithMeta(n int, metaData string) string {
	bytes := []byte(metaData)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 随机大小写英文字母字符串
func RandStr(n int) string {
	if n == 0 {
		n = RandInt(3, 20)
	}
	return RandStrWithMeta(n, AsciiLitter)
}

// 随机任意字节数组
func RandBytes(n int) []byte {
	if n == 0 {
		n = RandInt(3, 20)
	}
	dst := make([]byte, 0)
	for i := 0; i < n; i++ {
		dst = append(dst, byte(randInt(0, 255)))
	}
	return dst
}

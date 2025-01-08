package helper

import (
	"crypto/rand"
	"math/big"
	"unicode/utf16"
	"unsafe"
)

func UTF16ToString(buffer *uint16, length int) string {
	data := unsafe.Slice(buffer, length)
	return string(utf16.Decode(data))
}

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		randomByte, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[randomByte.Int64()]
	}
	return string(result), nil
}

package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func EncodeToBase63(originalUrl string) (string, error) {
	hash := sha256.Sum256([]byte(originalUrl))
	shortBytes := make([]byte, base64.RawURLEncoding.EncodedLen(len(hash)))
	base64.RawURLEncoding.Encode(shortBytes, hash[:])
	for i := range shortBytes {
		if !(shortBytes[i] >= 'A' && shortBytes[i] <= 'Z' ||
			shortBytes[i] >= 'a' && shortBytes[i] <= 'z' ||
			shortBytes[i] >= '0' && shortBytes[i] <= '9' ||
			shortBytes[i] == '_') {
			shortBytes[i] = '_'
		}
	}
	return string(shortBytes[:10]), nil
}

package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func CreateHexEncodePayload(payload []byte) string {
	hash := sha256.New()
	hash.Write(payload)
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateSymmetricSignature(key, data string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

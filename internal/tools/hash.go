package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacHash(data []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encode(str string) string {
	b := getSHA256Binary(str)
	h := hex.EncodeToString(b)
	return h
}

func getSHA256Binary(str string) []byte {
	r := sha256.Sum256([]byte(str))
	return r[:]
}

package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

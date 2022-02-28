package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func ShaEncrypt(data int64) string {
	shaCtx := sha256.New()
	str := strconv.FormatInt(data, 10)
	shaCtx.Write([]byte(str))
	cipherStr := shaCtx.Sum(nil)
	encryptedData := hex.EncodeToString(cipherStr)
	return encryptedData
}

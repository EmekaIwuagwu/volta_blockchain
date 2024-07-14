package utils

import (
    "crypto/sha256"
    "encoding/hex"
    "math/rand"
    "time"
)

func GenerateHex(n int) string {
    rand.Seed(time.Now().UnixNano())
    bytes := make([]byte, n)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}

func GenerateHash(input string) string {
    hash := sha256.New()
    hash.Write([]byte(input))
    return hex.EncodeToString(hash.Sum(nil))
}

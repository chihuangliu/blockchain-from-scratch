package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func Sha256Hash() string {
	rand.Seed(time.Now().UnixNano())

	randomBytes := make([]byte, 32)
	rand.Read(randomBytes)

	hashByte := sha256.Sum256(randomBytes)
	hashDecodedString := hex.EncodeToString(hashByte[:])

	return hashDecodedString

}

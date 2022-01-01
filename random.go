package tools

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// NewRandomString creates a new random string
func NewRandomString(length int, specials bool) string {
	if length < 1 {
		length = 10
	}

	// generate a random array of bytes and then use its int transposition as a seed
	b := make([]byte, 10)
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("Cannot seed math/rand package with cryptographically secure random number generator")
	}
	seededRand := rand.New(rand.NewSource(int64(binary.BigEndian.Uint64(b))))

	var charset string
	if specials {
		charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789" + "~=+%^*/()[]{}/!@#$?|"
	} else {
		charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	}

	buf := make([]byte, length)
	for i := range buf {
		buf[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(buf)
}

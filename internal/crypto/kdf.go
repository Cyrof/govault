package crypto

import (
	"bytes"
	"golang.org/x/crypto/argon2"
)

func KDF(password string, salt []byte) []byte {
	// 1 iteration, 64MB memory, 4 threads, 32-byte key
	return argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
}

func VerifyHash(password string, storedSalt, storedHash []byte) ([]byte, bool) {
	key := KDF(password, storedSalt)
	return key, bytes.Equal(key, storedHash)
}

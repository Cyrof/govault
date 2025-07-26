package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func (c *Crypto) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.Key)
	if err != nil {
		return nil, err
	}

	// use gcm mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// create nonce since gcm requires it
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// use seal command to encrypt using gcm
	ciphertext := aesGCM.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

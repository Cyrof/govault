package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type EncryptOptions struct {
	Key []byte
}

func (c *Crypto) Encrypt(data []byte, opt *EncryptOptions) ([]byte, error) {
	key := c.Key
	if opt != nil && opt.Key != nil {
		key = opt.Key
	}

	block, err := aes.NewCipher(key)
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

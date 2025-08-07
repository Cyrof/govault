package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type DecryptOptions struct {
	Key []byte
}

func (c *Crypto) Decrypt(ciphertext []byte, opt *DecryptOptions) ([]byte, error) {
	key := c.Key
	if opt != nil && opt.Key != nil {
		key = opt.Key
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, data := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return aesGCM.Open(nil, nonce, data, nil)
}

package crypto

import (
	"errors"
)

func (c *Crypto) SetupNewPassword(password string) error {
	salt, _ := GenerateSalt(32)
	key := KDF(password, salt)
	c.Salt = salt
	c.Key = key
	return nil
}

func (c *Crypto) SetupFromMeta(password string, storedSalt, storedHash []byte) error {
	key, ok := VerifyHash(password, storedSalt, storedHash)
	if !ok {
		return errors.New("invalid password")
	}

	c.Salt = storedSalt
	c.Key = key
	return nil
}

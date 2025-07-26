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
	key := KDF(password, storedSalt)
	if !VerifyHash(key, storedHash) {
		return errors.New("Invalid Password")
	}

	c.Salt = storedSalt
	c.Key = key
	return nil
}

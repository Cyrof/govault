package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func GeneratePassword(opts Options) (string, error) {
	charsetStr := BuildCharset(opts)
	if charsetStr == "" {
		return "", errors.New("no charset sets selected")
	}
	password := make([]byte, opts.Length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charsetStr))))
		if err != nil {
			return "", err
		}
		password[i] = charsetStr[n.Int64()]
	}

	return string(password), nil
}

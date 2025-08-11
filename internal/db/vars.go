package db

import "errors"

var ErrDuplicateKey = errors.New("secret with this key already exists")

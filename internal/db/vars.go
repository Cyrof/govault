package db

import "errors"

var ErrDuplicateKey = errors.New("secret with this key already exists")

var ErrNotFound = errors.New("secret not found")

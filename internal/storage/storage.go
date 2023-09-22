package storage

import "errors"

var (
	// ErrURLNotFound is error for wrong URL
	ErrURLNotFound = errors.New("url not found")
	// ErrURLExists is error if record with URL already exists
	ErrURLExists = errors.New("url exists")
)

package storage

import "errors"

var (
	// Error for wrong URL
	ErrURLNotFound = errors.New("url not found")
	// Error if record with URL already exists
	ErrURLExists = errors.New("url exists")
)

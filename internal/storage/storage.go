package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLNotExist = errors.New("url does not exist")
	ErrURLExists   = errors.New("url already exists")
)

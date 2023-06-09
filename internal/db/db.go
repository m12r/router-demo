package db

import "errors"

type DB interface {
	Get(key string) (string, error)
	Put(key string, value string) error
	Delete(key string) error
}

var (
	ErrNotFound = errors.New("not found")
)

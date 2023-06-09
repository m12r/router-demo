package mapdb

import "github.com/m12r/router-demo/internal/db"

type DB map[string]string

func (d DB) Get(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", db.ErrNotFound
	}
	return value, nil
}

func (d DB) Put(key string, value string) error {
	d[key] = value
	return nil
}

func (d DB) Delete(key string) error {
	delete(d, key)
	return nil
}

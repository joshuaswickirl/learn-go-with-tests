package maps

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", notFoundError{}
	}
	return v, nil
}

type notFoundError struct{}

func (nfe notFoundError) Error() string {
	return "could not find the word you were looking for"
}

func (nfe notFoundError) NotFound() bool {
	return true
}

func (d Dictionary) Add(k, v string) error {
	if _, ok := d[k]; ok {
		return fmt.Errorf("value already exists for given key %q", k)
	}
	d[k] = v
	return nil
}

func (d Dictionary) Update(k, v string) error {
	if _, ok := d[k]; !ok {
		return errors.New("cannot update records that don't exist")
	}
	d[k] = v
	return nil
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}

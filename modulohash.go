// Package modulohash maps a string into a list using CRC32.
package modulohash

import (
	"errors"
	"fmt"
	"hash/crc32"
)

// ModuloHash is the list of targets.
type ModuloHash struct {
	targets []string
}

// NewModuloHash will create a new ModuloHash.
func NewModuloHash(targets []string) (*ModuloHash, error) {
	if len(targets) == 0 {
		return nil, errors.New("no targets")
	}
	h := new(ModuloHash)
	h.targets = make([]string, len(targets))
	copy(h.targets, targets)
	return h, nil
}

// New is an alias to NewModuloHash.
func New(targets []string) (*ModuloHash, error) {
	return NewModuloHash(targets)
}

// Find will fairly determine a target value in the array Targets by using the crc32 hash value
// mod'd with the number of targets in the target list.
func (h *ModuloHash) Find(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty source string")
	}
	l := len(h.targets)
	sSum := crc32.ChecksumIEEE([]byte(s))
	i := sSum % uint32(l)
	if int(i) > (l - 1) {
		e := fmt.Sprintf("%d is out of bounds", i)
		return "", errors.New(e)
	}
	return h.targets[i], nil
}

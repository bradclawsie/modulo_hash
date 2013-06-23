package modulo_hash

import (
	"fmt"
	"errors"
	"hash/crc32"
)

// The ModuloHash structure is just the list of targets. 
type ModuloHash struct {
	Targets []string
}

// Create a new ModuloHash
func NewModuloHash() (*ModuloHash) {
	mh := new(ModuloHash)
	mh.Targets = make([]string,0)
	return mh
}

// fairly find a target value in the array Targets by using the crc32 hash value 
// mod'd with the number of targets in the target list
func (h *ModuloHash) Find(s string) (string,error) {
	if len(h.Targets) == 0 {
		return "",errors.New("cannot match to empty targets list")
	}
	if s == "" {
		return "",errors.New("cannot match empty source string")
	}
	s_sum := crc32.ChecksumIEEE([]byte(s))
	l := len(h.Targets) 
	i := s_sum % uint32(l)
	if int(i) > (len(h.Targets)-1) {
		e := fmt.Sprintf("%d is out of bounds",i)
		return "",errors.New(e)
	}
	return h.Targets[i],nil
}

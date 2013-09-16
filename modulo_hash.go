// A package for fairly hashing keys to a small set of values. Uses crc32 to create uniformly distributed values that
// are then hashed again to a target list of your choice using modulo indexing. Useful for work distribution, etc.
package modulo_hash

import (
	"fmt"
	"errors"
	"hash/crc32"
	"sync"
)

// The ModuloHash structure is just the list of targets. 
type ModuloHash struct {
	targets []string
	mutex sync.RWMutex
}

// NewModuloHash will create a new ModuloHash.
func NewModuloHash() (*ModuloHash) {
	h := new(ModuloHash)
	h.targets = make([]string,0)
	return h
}

// New is an alias to NewModuloHash.
func New() (*ModuloHash) {
	return NewModuloHash()
}

// GetTargets will read a copy of the targets list.
func (h *ModuloHash) GetTargets() []string {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	tgts := make([]string,len(h.targets))
	copy(tgts,h.targets)
	return tgts
}

// SetTargets will write a new targets list.
func (h *ModuloHash) SetTargets(tgts []string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.targets = make([]string,len(tgts))
	copy(h.targets,tgts)
}

// Find will fairly determine a target value in the array Targets by using the crc32 hash value 
// mod'd with the number of targets in the target list.
func (h *ModuloHash) Find(s string) (string,error) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	l := len(h.targets) 
	if l == 0 {
		return "",errors.New("cannot match to empty targets list")
	}
	if s == "" {
		return "",errors.New("cannot match empty source string")
	}
	s_sum := crc32.ChecksumIEEE([]byte(s))
	i := s_sum % uint32(l)
	if int(i) > (l-1) {
		e := fmt.Sprintf("%d is out of bounds",i)
		return "",errors.New(e)
	}
	return h.targets[i],nil
}

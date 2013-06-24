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

// Create a new ModuloHash
func NewModuloHash() (*ModuloHash) {
	h := new(ModuloHash)
	h.targets = make([]string,0)
	return h
}

// Read a copy of the targets list
func (h *ModuloHash) GetTargets() []string {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	tgts := make([]string,len(h.targets))
	copy(tgts,h.targets)
	return tgts
}

// Write a new targets list
func (h *ModuloHash) SetTargets(tgts []string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.targets = make([]string,len(tgts))
	copy(h.targets,tgts)
}

// Fairly find a target value in the array Targets by using the crc32 hash value 
// mod'd with the number of targets in the target list
func (h *ModuloHash) Find(s string) (string,error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
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

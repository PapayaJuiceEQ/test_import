package db

import "sync"

var (
	i int = 0
	m sync.Mutex
)

// IncI ...
func IncI() {
	m.Lock()
	i++
	m.Unlock()
}

// GetI ...
func GetI() int {
	m.Lock()
	ii := i
	m.Unlock()
	return ii
}

package counter

import "sync"

type mem struct {
	sync.Mutex
	count uint64
}

func (m *mem) Count() (uint64, error) {
	m.Lock()
	m.count++
	c := m.count
	m.Unlock()
	return c, nil
}

func newMem() (*mem, error) { return &mem{}, nil }

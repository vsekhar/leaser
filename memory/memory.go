// Package memory implements a leaser driver in memory (for testing).
package memory

import (
	"sync"
	"time"

	"github.com/vsekhar/leaser/driver"
)

type state struct {
	expiry time.Time
	value  string
}

func (s *state) Expiry() time.Time {
	return s.expiry
}

func (s *state) Value() string {
	return s.value
}

type memleaser struct {
	expiry time.Time
	value  string
	mu     sync.Mutex
}

func (m *memleaser) Get() (driver.State, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return &state{m.expiry, m.value}, nil
}

func (m *memleaser) Set(s driver.State, v string, expiry time.Time) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	ms, ok := s.(*state)
	if !ok {
		panic("bad state")
	}
	if !m.expiry.Equal(ms.expiry) || m.value != ms.value {
		return driver.ErrStateChanged
	}
	m.expiry = expiry
	m.value = v
	return nil
}

// New returns a new memory leaser driver.
func New() driver.Interface {
	return &memleaser{}
}

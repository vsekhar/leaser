// Package leaser provides atomic lease semantics on top of any backing store
// that can support compare-and-set (CAS).
package leaser

import (
	"time"

	"github.com/vsekhar/leaser/driver"
)

// Leaser allows clients to obtain leases on behalf of values.
//
// Leaser can typically be used to acquire a lease on behalf a client identified
// by its host name.
type Leaser struct {
	d   driver.Interface
	s   driver.State
	cfg Config
}

func (l *Leaser) isPast(t time.Time) bool {
	return t.Before(time.Now().Add(l.cfg.MaxClockSkew))

}

// Acquire attempts to obtain a lease for value v. Acquire returns the currently
// valid value and expiry or an error.
func (l *Leaser) Acquire(v string, expiry time.Time) (string, time.Time, error) {
	var err error
	l.s, err = l.d.Get()
	if err != nil {
		return "", time.Time{}, err
	}
	if l.isPast(l.s.Expiry()) {
		// write new value
		err = l.d.Set(l.s, v, expiry)
		if err != nil {
			return "", time.Time{}, err
		}
		return v, expiry, nil
	}
	// return old (still valid) value
	return l.s.Value(), l.s.Expiry(), nil
}

// New returns a Leaser using driver d.
func New(d driver.Interface, cfg Config) *Leaser {
	return &Leaser{d: d, cfg: cfg}
}

// Config is a configuration for a Leaser.
type Config struct {
	MaxClockSkew time.Duration
}

// DefaultConfig is a default configuration for a Leaser.
var DefaultConfig Config = Config{
	MaxClockSkew: 10 * time.Millisecond,
}

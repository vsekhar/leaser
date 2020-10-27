package driver

import (
	"fmt"
	"time"
)

// Interface is the interface a leaser driver must provide.
type Interface interface {
	// Get gets the current state of the leaser.
	Get() (State, error)

	// Set atomically writes value to the leaser if and only if the leaser
	// matches state s at the time of the write.
	Set(s State, value string, expiry time.Time) error
}

// State is the state of the lease as obtained by Get. It should contain any
// unexported fields needed to ensure atomic operation between calls to Get and
// Set.
type State interface {
	Expiry() time.Time
	Value() string
}

// ErrStateChanged is returned by Set if the underlying state of the leaser
// changed between calls to Get and Set.
var ErrStateChanged = fmt.Errorf("state has changed between calls to Get and Set")

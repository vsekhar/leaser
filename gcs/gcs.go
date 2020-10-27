// Package gcs provides a leaser driver backed by Google Cloud Storage.
package gcs

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
)

const attrPrefix = "__golang_leaser__"

type state struct {
	expiry time.Time
	value  string

	metaGeneration int64
}

func (s *state) Expiry() time.Time {
	return s.expiry
}

func (s *state) Value() string {
	return s.value
}

type driver struct {
	obj *storage.ObjectHandle
}

func (d *driver) Get() (driver.State, error) {
	ifobj := d.obj
	if d.metaGeneration 
	attrs, err := d.obj.Attrs(context.Background())
	if err == storage.ErrObjectNotExist {
		// no lease
		return &state{expiry: time.Time{}, value: "", }
	}
	if err != nil {
		return nil, err
	}
}

// New returns a new leaser driver backed by the GCS object obj.
func New(obj *storage.ObjectHandle) driver.Interface {
	return &driver{
		obj: obj,
	}
}

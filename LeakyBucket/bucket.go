package LeakyBucket

import (
	"errors"
	"time"
)

var (
	// ErrorFull is returned when the amount requested to add exceeds the remaining space in the bucket.
	ErrorFull = errors.New("add exceeds free capacity")
)

type BucketI interface {
	Capacity() uint
	Remaining() uint
	Reset() time.Time
	Add(uint) (BucketState, error)
}

type BucketState struct {
	Capacity uint
	Remaining uint
	Reset time.Time
}

type StorageI interface {
	Create(name string, capacity uint, rate time.Duration) (BucketI, error)
}




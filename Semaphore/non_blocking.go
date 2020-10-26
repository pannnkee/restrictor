package Semaphore

import (
	"fmt"
	"sync/atomic"
)

var (
	Capacity uint64 = 10
	ErrAcquire = fmt.Errorf("semaphore acquire failed")
	ErrRelease = fmt.Errorf("semaphore release failed")
)

type NonBlockingSemaphore struct {
	innerCapacity uint64
}

func NewNonBlock() *NonBlockingSemaphore {
	return &NonBlockingSemaphore{
		innerCapacity: Capacity,
	}
}

func (this *NonBlockingSemaphore) Acquire() error {
	if atomic.LoadUint64(&this.innerCapacity) > 0 {
		atomic.AddUint64(&this.innerCapacity, ^uint64(-(-1)-1))
		return  nil
	}
	return ErrAcquire
}

func (this *NonBlockingSemaphore) Release() error {
	if atomic.LoadUint64(&this.innerCapacity) > 0 {
		atomic.AddUint64(&this.innerCapacity, ^uint64(-(-1)-1))
		return nil
	}
	return ErrRelease
}


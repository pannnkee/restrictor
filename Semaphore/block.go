package Semaphore

import (
	"fmt"
	"time"
)

type BlockSemaphore struct {
	innerChan chan struct{}
}

var (
	ErrSemaphoreBlocking = fmt.Errorf("semaphore acquire is blocking")
)

func NewBlock(capacity uint) *BlockSemaphore {
	return &BlockSemaphore{
		innerChan: make(chan struct{}, capacity), 
	}
}

func (this *BlockSemaphore) Acquire() error {
	select {
	case <-this.innerChan:
		return nil
	default:
		time.Sleep(time.Millisecond*100)
		return ErrSemaphoreBlocking
	}
}

func (this *BlockSemaphore) Release() error {
	select {
	case this.innerChan<- struct{}{}:
		return nil
	default:
		return ErrSemaphoreBlocking
	}
}


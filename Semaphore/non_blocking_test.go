package Semaphore

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNonBlock(t *testing.T) {
	convey.Convey("start testing", t, func() {
		convey.Convey("New", func() {
			semaphore := NewNonBlock()
			if semaphore == nil {
				t.Errorf("new NonBlock failed, semaphore is nil")
			}
		})

		convey.Convey("Acquire", func() {
			semaphore := NewNonBlock()
			for i := 0; uint64(i) < Capacity; i++ {
				err := semaphore.Acquire()
				if err != nil {
					t.Errorf("Acquire failed")
				}
			}

			err2 := semaphore.Acquire()
			if err2 == nil {
				t.Errorf("over size")
			}
		})
	})
}

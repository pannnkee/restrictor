package Semaphore

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestBlock(t *testing.T) {
	convey.Convey("starting test", t, func() {
		convey.Convey("New", func() {
			semaphore := NewBlock(10)
			if semaphore == nil {
				t.Fatalf("New BlockSemaphore error")
			}
			if cap(semaphore.innerChan) != 10 {
				t.Fatalf("New BlockSemaphore innerChan cap != capacity")
			}
		})

		convey.Convey("Acquire", func() {
			semaphore := NewBlock(10)
			err := semaphore.Acquire()
			if err == nil {
				println("Acquire success")
			} else {
				println(err)
			}
		})

		convey.Convey("Release", func() {
			semaphore := NewBlock(10)
			err := semaphore.Release()
			if err == nil {
				println("Release success")
			} else {
				println(err)
			}
		})

		convey.Convey("Run", func() {
			semaphore := NewBlock(10)

			go func() {
				for {
					semaphore.Acquire()
					time.Sleep(time.Millisecond * 500)
				}
			}()

			go func() {
				for  {
					semaphore.Release()
					time.Sleep(time.Second)
				}
			}()
		})
	})
}
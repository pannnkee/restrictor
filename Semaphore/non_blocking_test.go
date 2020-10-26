package Semaphore

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
	"time"
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

		convey.Convey(".", func() {
			semaphore := NewNonBlock()
			for i:=0; i<100; i++ {
				go func() {
					err := semaphore.Acquire()
					if err == nil {
						//do
						rand.Seed(time.Now().Unix())
						time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
						err := semaphore.Release()
						if err != nil {
							t.Errorf(ErrRelease.Error())
						}
					} else {
						fmt.Println("err:", err)
					}
				}()
			}
		})
	})
}

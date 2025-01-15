package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {

	t.Run("merge channels", func(t *testing.T) {
		ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
		go func() {
			for i := 0; i < 15; i++ {
				ch1 <- i
			}
			close(ch1)
		}()
		go func() {
			for i := 0; i < 10; i++ {
				ch2 <- i
			}
			close(ch2)
		}()
		go func() {
			for i := 0; i < 5; i++ {
				ch3 <- i
			}
			close(ch3)
		}()

		for out := range mergeChannels(ch1, ch2, ch3) {
			t.Log(out)
		}
	})

}

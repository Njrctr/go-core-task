package main

import (
	"testing"
)

func TestGetRandomIntNumber(t *testing.T) {
	t.Run("GetRandomIntNumber", func(t *testing.T) {
		ch := make(chan int)
		go getRandomIntNumber(ch)
		actual := <-ch
		if actual == 0 {
			t.Errorf("expected any int, got %v", actual)
		}
	})
}

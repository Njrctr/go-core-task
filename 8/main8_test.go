package main

import "testing"

func TestWaitGroup(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		wg := NewWaitGroup(100)
		wg.Add(10)
		if wg.counter != 10 {
			t.Error("After wg.Add(10) counter is not 10")
		}
		wg.Done()
		if wg.counter != 9 {
			t.Error("After wg.Done() counter is not 9")
		}
	})
}

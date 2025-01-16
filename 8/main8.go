package main

import (
	"fmt"
	"time"
)

type WaitGroup struct {
	sem     chan struct{}
	counter int
}

func NewWaitGroup(max int) *WaitGroup {
	return &WaitGroup{
		sem: make(chan struct{}, max),
	}
}

func (wg *WaitGroup) Add(count int) {
	wg.counter += count
	for i := 0; i < count; i++ {
		wg.sem <- struct{}{}
	}
}

func (wg *WaitGroup) Done() {
	if wg.counter > 0 {
		wg.counter--
		<-wg.sem
	}
}

func (wg *WaitGroup) Wait() {
	for wg.counter > 0 {
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	wg := NewWaitGroup(100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("Cogoroutine %d is done!\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("All gorutines are done!")
}

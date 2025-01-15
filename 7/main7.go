package main

import "sync"

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

	go func() {
		for i := 0; i < 12; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 100; i < 135; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	go func() {
		for i := 1000; i < 1077; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	for out := range mergeChannels(ch1, ch2, ch3) {
		println(out)
	}
}

func mergeChannels(channels ...<-chan int) <-chan int {
	resultChannel := make(chan int)

	wg := &sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for val := range ch {
				resultChannel <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	return resultChannel

}

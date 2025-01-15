package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- uint8(i)
			time.Sleep(3 * time.Second)
		}
		close(ch1)
	}()

	go conveer(ch1, ch2)
	for val := range ch2 {
		fmt.Println(val)
	}
}

func conveer(ch1 chan uint8, ch2 chan float64) {
	defer close(ch2)
	for {
		select {
		case val, open := <-ch1:
			if !open {
				return
			}
			floatNum := float64(val)
			ch2 <- math.Pow(floatNum, 2)
		default:
			continue
		}

	}
}

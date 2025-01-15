package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go getRandomIntNumber(ch)
	for {
		fmt.Println(<-ch)
		time.Sleep(1 * time.Second)
	}

}

func getRandomIntNumber(ch chan int) {
	for {
		ch <- rand.Int()
	}
}

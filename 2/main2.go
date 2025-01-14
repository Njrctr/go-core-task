package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	var randomSlice []int
	for i := 0; i < 10; i++ {
		randomSlice = append(randomSlice, rand.Intn(1000))
	}
	fmt.Println("randomSlice:", randomSlice)
	newSlice := sliceExample(randomSlice)
	fmt.Println("newSlice:", newSlice)
	newAfterAppend := addElements(newSlice, 345345)
	fmt.Println("newAfterAppend:", newAfterAppend)
	sliceCopy := copySlice(newAfterAppend)
	fmt.Println("sliceCopy:", sliceCopy)
	newAfterRemove, _ := removeElement(sliceCopy, 2)
	fmt.Println("newAfterRemove:", newAfterRemove)
}

func sliceExample(old []int) []int {
	var newSlice = []int{}

	for _, el := range old {
		if el%2 == 0 {
			newSlice = append(newSlice, el)
		}
	}
	return newSlice
}

func addElements(old []int, newElement int) []int {
	return append(old, newElement)
}

func copySlice(old []int) []int {
	if len(old) == 0 {
		return []int{}
	}
	newSlice := make([]int, len(old))
	copy(newSlice, old)
	return newSlice
}

func removeElement(old []int, index int) ([]int, error) {
	if index >= len(old) {
		return nil, errors.New("index out of range")
	}
	return append(old[:index], old[index+1:]...), nil
}

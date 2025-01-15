package main

import "fmt"

func main() {
	a := []int{64, 2, 3, 43}
	b := []int{64, 2, 3, 43}
	hasIntersection, intersection := intersectionOfSlices(a, b)
	fmt.Println(hasIntersection, intersection)
}

func intersectionOfSlices(a, b []int) (bool, []int) {
	if len(a) == 0 || len(b) == 0 {
		return false, []int{}
	}
	intersectionSlice := []int{}
	intersectionMap := make(map[int]int)
	for _, v := range a {
		intersectionMap[v] = 1
	}
	for _, v := range b {
		if _, ok := intersectionMap[v]; ok {
			intersectionSlice = append(intersectionSlice, v)
		}
	}

	if len(intersectionSlice) == 0 {
		return false, intersectionSlice
	}

	return true, intersectionSlice
}

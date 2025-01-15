package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"b", "c", "d"}
	uniqueSlice := getUniqueBySlices(a, b)
	fmt.Println(uniqueSlice)
}

func getUniqueBySlices(a []string, b []string) []string {
	uniqueMap := map[string]int{}
	for _, value := range a {
		uniqueMap[value] = 1
	}
	for _, value := range b {
		delete(uniqueMap, value)
	}
	uniqueSlice := []string{}
	for key := range uniqueMap {
		uniqueSlice = append(uniqueSlice, key)
	}

	return uniqueSlice
}

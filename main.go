package main

import "fmt"

func main() {
	needle := 6
	haystack := []int{1, 3, 4, 6, 7, 9, 10, 11, 13}

	fmt.Printf("Looking for: %v in this haystack: %v\n", needle, haystack)

	index := binarySearch(haystack, needle)
	if index >= 0 {
		fmt.Printf("Found the needle: %v at index: %v\n", needle, index)
	} else {
		fmt.Printf("Did not find the needle %v:\n", needle)
	}
}

func binarySearch(haystack []int, needle int) int {
	lo := 0
	hi := len(haystack) - 1

	midIndex := lo + (hi-lo)/2
	midValue := haystack[midIndex]

	if midValue == needle {
		return midIndex
	} else if midValue > needle {

	} else {

	}

	return -1
}

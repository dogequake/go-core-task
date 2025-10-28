package main

import (
	"fmt"
	"math/rand"
)

var originalSlice [10]int

func main() {
	fillingSlice()
	fmt.Println(originalSlice)
	modifiedSlice := sliceExample(originalSlice[:])
	fmt.Println(modifiedSlice)
	extendedSlice := addElements(modifiedSlice, 5)
	fmt.Println(extendedSlice)
	copiedSlice := copySlice(extendedSlice)
	fmt.Println(copiedSlice)
	finalSlice := removeElement(copiedSlice, 2)
	fmt.Println(finalSlice)
}

func fillingSlice() {
	for i := 0; i < len(originalSlice); i++ {
		originalSlice[i] = rand.Intn(100)
	}
}

func sliceExample(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(slice []int, element int) []int {
	return append(slice, element)
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

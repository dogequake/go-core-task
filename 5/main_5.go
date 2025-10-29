package main

import "fmt"

var a = []int{65, 3, 58, 678, 64}
var b = []int{64, 2, 3, 43}

func main() {
	ok, result := intersectSlices(a, b)
	fmt.Printf("%t, %#v\n", ok, result)
}

func intersectSlices(slice1, slice2 []int) (bool, []int) {
	newSlice := make([]int, 0)

	for _, s2 := range slice2 {
		for _, s1 := range slice1 {
			if s1 == s2 {
				newSlice = append(newSlice, s1)
				break
			}
		}
	}

	return len(newSlice) > 0, newSlice
}

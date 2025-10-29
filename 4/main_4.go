package main

import (
	"fmt"
)

var slice1 = []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
var slice2 = []string{"banana", "date", "fig"}

func main() {
	newSlice := sortSlice(slice1, slice2)
	fmt.Println(newSlice)
}

func sortSlice(slice1, slice2 []string) []string {
	result := make([]string, 0)

	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		if !found {
			result = append(result, s1)
		}
	}
	return result
}

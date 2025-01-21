package utils

import "fmt"

func init() {
	fmt.Println("Function init in package utils")
}

func InsSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}

func InSliceInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}

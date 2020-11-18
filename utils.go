package main

import (
	"strings"
	"unicode"
)

func trimSpace(str string) string {
	// Trims all spaces from string
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func makeRange(min int, max int) []int {
	// Returns a slice of int within range
	ret := make([]int, max-min+1)
	for i := range ret {
		ret[i] = min + i
	}
	return ret
}

//TODO Generalize *InSlice funcs (1 function)
func findStrInSlice(slice []string, val string) bool {
	// returns true if value in slice
	for _, can := range slice {
		if can == val {
			return true
		}
	}
	return false
}

//func findIntInSlice(slice []int, val int) bool {
//	// returns true if value in slice
//	for _, can := range slice {
//		if can == val {
//			return true
//		}
//	}
//	return false
//}

func removeElmIntInSlice(slice []int, i int) []int {
	// Removes Element(INT) from slice, RETURNS UNORDERED
	slice[len(slice)-1], slice[1] = slice[i], slice[len(slice)-1]
	return slice[:len(slice)-1]
}

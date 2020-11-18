package main

import (
	"strings"
)

const (
	MIN = 97
	MAX = 122
)

func encode(str string, offset int) string {
	// Encodes string using a Caesar Cipher.
	// Can be used to decode using the inverse (i.e. +2 becomes -2)
	orgStr := strings.Split(strings.ToLower(str), " ")
	retString := ""
	for _, str := range orgStr {
		strBuild := ""
		for _, chr := range []rune(str) {
			num := shiftChar(int(chr), offset)
			strBuild = strBuild + string(num)
		}
		retString = retString + strBuild
	}
	return retString
}

func shiftChar(org int, off int) int {
	// Loop int between a MIN / MAX value
	if org > MAX || org < MIN {
		return org
	}
	switch val := org + off; {

	case val > MAX:
		return val - 26

	case val < MIN:
		return val + 26

	default:
		return val
	}
}

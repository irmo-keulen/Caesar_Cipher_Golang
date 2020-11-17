package main

import (
	"errors"
	"fmt"
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
			num, err := checkWrap(int(chr), offset)
			if err != nil {
				fmt.Println(err)
			}
			strBuild = strBuild + string(num)
		}
		retString = retString + " " + strBuild
	}
	return retString
}

func checkWrap(org int, off int) (int, error) {
	// Loop int between a MIN / MAX value
	if org > MAX || org < MIN {
		return 0, errors.New("Character out of bound : Value Error")
	}
	val := org + off

	if val < MIN {
		return MAX - (MIN - val), nil
	} else if val > MAX {
		return MIN + (val - MAX), nil
	}
	return val, nil
}

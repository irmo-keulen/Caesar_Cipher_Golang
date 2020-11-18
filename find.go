package main

import (
	"io"
	"os"
	"strings"
)

const bufferSize = 100

func search(path string, word string) ([]int, error) {
	// TODO Refactor function name
	var retInt []int
	length := len(word)
	file, err := os.Open(path)

	if err != nil {
		return retInt, err
	}
	defer file.Close()

	buffer := make([]byte, bufferSize)
	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				return retInt, err
			}
			break
		}
		words := strings.Split(string(buffer[:bytesread]), " ")
		for _, wrd := range words {
			if len(wrd) == length {
				for i := 0; i <= 27; i++ {
					if word == encode(wrd, i) {
						retInt = append(retInt, i)
					}
				}
			}
		}
	}
	return rmDuplicate(retInt), nil
}

func rmDuplicate(intSlice []int) []int {
	k := make(map[int]bool)
	list := []int{}
	for _, val := range intSlice {
		if _, value := k[val]; !value {
			k[val] = true
			list = append(list, val)
		}
	}
	return list
}

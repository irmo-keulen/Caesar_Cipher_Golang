package main

import (
	"io"
	"os"
	"strings"
)

func searchInFile(path string, word string, bufferSize int) ([]int, error) {
	// TODO Refactor function name
	var retInt []int
	var foundWords []string
	possibilities := makeRange(0, 27)
	length := len(word)

	if bufferSize == 0 {
		// bufferSize can be set using flag
		bufferSize = 100
	}

	file, err := os.Open(path)
	if err != nil {
		return retInt, err
	}
	defer file.Close()

	buffer := make([]byte, bufferSize)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				return retInt, err
			}
			break
		}
		words := strings.Split(string(buffer[:bytesRead]), " ")

		// TODO Add Go routine & optimize
		for _, wrd := range words {
			if len(wrd) == length && !findStrInSlice(foundWords, wrd) {
				foundWords = append(foundWords, wrd)
				for idx := range possibilities {
					if word == encode(wrd, idx) {
						retInt = append(retInt, idx)
						removeElmIntInSlice(possibilities, idx)
					}
				}
			}
		}
	}
	return retInt, nil
}

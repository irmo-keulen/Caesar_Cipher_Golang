package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	offsetFlag := flag.Int("n", 0, "Number of shifts on the string")
	capsFlag := flag.Bool("c", false, "Output string in caps")
	trimFlag := flag.Bool("t", false, "removes all whitespace")
	pathFlag := flag.String("p", "", "The absolute path of the file that will be searched")
	flag.Parse()

	str := encode(strings.Join(flag.Args(), " "), *offsetFlag)
	caps, trim, path := *capsFlag, *trimFlag, *pathFlag

	if trim {
		str = trimWhiteSpace(str)
	}
	if caps {
		str = strings.ToUpper(str)
	}
	fmt.Println(str)
	output, _ := search(path, "hello")
	fmt.Println(output)
}

func trimWhiteSpace(str string) string {
	// Trims all whitespace from string
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

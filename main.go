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
	flag.Parse()

	str := encode(strings.Join(flag.Args(), " "), *offsetFlag)
	caps, trim := *capsFlag, *trimFlag

	if trim {
		str = trimWhiteSpace(str)
	}

	if caps {
		fmt.Println(strings.ToUpper(str))
	} else {
		fmt.Println(str)
	}
}

func trimWhiteSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

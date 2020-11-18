package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	offsetFlag := flag.Int("n", 0, "Number of shifts on the string")
	capsFlag := flag.Bool("c", false, "Output string in caps")
	trimFlag := flag.Bool("t", false, "removes all whitespace")
	pathFlag := flag.String("p", "", "The absolute path of the file that will be searched")
	bufferFlag := flag.Int("b", 0, "BufferSize (higher uses more Memory) default = 100")
	flag.Parse()

	str := encode(strings.Join(flag.Args(), " "), *offsetFlag)
	caps, trim, path, buffer := *capsFlag, *trimFlag, *pathFlag, *bufferFlag

	if trim {
		str = trimSpace(str)
	}
	if caps {
		str = strings.ToUpper(str)
	}
	fmt.Println(str)
	output, _ := searchInFile(path, "hello", buffer)
	fmt.Println(output)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(encode(strings.Join(os.Args[1:], " "), 1))

}

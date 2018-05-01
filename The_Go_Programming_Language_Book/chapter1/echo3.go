package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// using joins from strings package
	fmt.Println(strings.Join(os.Args[1:], " "))
}

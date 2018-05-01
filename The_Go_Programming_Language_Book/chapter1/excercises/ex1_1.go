// Modify the program so that it prints os.Args[0], it's own name
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// using joins from strings package
	fmt.Println(strings.Join(os.Args[0:], " "))
}

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	firstMethod()
	firstMethod()
	firstMethod()
	result1 := time.Since(start1).Nanoseconds()
	fmt.Println(result1)
	start2 := time.Now()
	secondMethod()
	secondMethod()
	secondMethod()
	result2 := time.Since(start2).Nanoseconds()
	fmt.Println(result2)
}

func firstMethod() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func secondMethod() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))

	start := time.Now()
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
	finish := time.Since(start)
	fmt.Println("Execution time: ns", finish.Nanoseconds())

	start1 := time.Now()
	result1 := strings.Join(os.Args[1:], "\n")
	fmt.Println(result1)
	finish1 := time.Since(start1)
	fmt.Println("Execution time: ns", finish1.Nanoseconds())

}

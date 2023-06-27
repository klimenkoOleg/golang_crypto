package main

import (
	"bytes"
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

	fmt.Println("Commas split: " + comma("1234567890"))

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {

	var buf bytes.Buffer

	//n := 0

	n := len(s) % 3
	for n <= 3 {
		buf.Write(s[n : n+3])
		n += 3
	}

	return buf.String()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your data")

	buf := ""
	for {
		s, _ := inputReader.ReadString('\n')
		if s == "S\n" {
			break
		}
		nrlines++
		buf += strings.TrimRight(s, "\n")
	}
	Counters(buf)
	fmt.Printf("The number of bytes including spaces: %v\n", nrchars)
	fmt.Printf("The number of words: %v\n", nrwords)
	fmt.Printf("The number of lines: %v\n", nrlines)
}

func Counters(input string) {
	nrchars = len([]byte(input))
	words := SplitAny(input, " \n")
	nrwords = len(words) + nrlines - 1
}

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

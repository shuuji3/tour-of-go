package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	words := strings.Split(s, " ")
	for _, word := range words {
		counter[word] += 1
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}

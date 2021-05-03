package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	for _, v := range strings.Fields(s) {
		count[v]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}

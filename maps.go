package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		m[w] += 1
	}
	fmt.Println(m)
	return m
}

func main() {
	wc.Test(WordCount)
}

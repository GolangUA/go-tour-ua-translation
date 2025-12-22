//go:build OMIT

package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// TODO: поверніть map, де ключі — слова з s, а значення — кількість появ.
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}

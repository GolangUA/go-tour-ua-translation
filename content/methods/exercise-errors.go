//go:build OMIT

package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	// TODO: реалізуйте Sqrt: для від’ємних x повертайте ErrNegativeSqrt.
	// Поверніть значення кореня та помилку (nil у випадку успіху).
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

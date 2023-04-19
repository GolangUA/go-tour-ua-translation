//go:build OMIT
// +build OMIT

package main

import "fmt"

const (
	// Створіть величезне число, зсунувши на 1 біт вліво на 100 позицій.
	// Іншими словами число 1, за яким слідує 100 нулів.
	Big = 1 << 100
	// Зсуваємо його знову вправо на 99 позицій, так що отримуємо 1<<1, або просто 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

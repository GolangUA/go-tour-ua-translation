//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Тип: %T Значення: %v\n", ToBe, ToBe)
	fmt.Printf("Тип: %T Значення: %v\n", MaxInt, MaxInt)
	fmt.Printf("Тип: %T Значення: %v\n", z, z)
}

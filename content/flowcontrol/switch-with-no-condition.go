//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Доброго ранку!")
	case t.Hour() < 17:
		fmt.Println("Добрий день.")
	default:
		fmt.Println("Добрий вечір.")
	}
}

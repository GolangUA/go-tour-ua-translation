//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Коли Субота?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Сьогодні.")
	case today + 1:
		fmt.Println("Завтра.")
	case today + 2:
		fmt.Println("Через два дні.")
	default:
		fmt.Println("Колись буде..")
	}
}

//go:build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// Виводимо привітання і поточний час.
	fmt.Println("Вітаємо в Playground!")

	fmt.Println("Час:", time.Now())
}

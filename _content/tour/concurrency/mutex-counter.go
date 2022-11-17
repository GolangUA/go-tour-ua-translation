//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter безпечний для одночасного використання.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc збільшує лічільник для ключа key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock в цьому місці забезпечує доступ до словника c.v лише однієї goroutine за раз.
	c.v[key]++
	c.mu.Unlock()
}

// Value повертає поточне значення лічильника для ключа key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock в цьому місці забезпечує доступ до словника c.v лише однієї goroutine за раз.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("якийсьключ")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("якийсьключ"))
}

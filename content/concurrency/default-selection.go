//go:build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] цок.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] БАЦ!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     чекання.\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

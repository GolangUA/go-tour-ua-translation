//go:build OMIT

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go працює на ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		// (та інші)
		fmt.Printf("%s.\n", os)
	}
}

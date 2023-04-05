// +build no-build OMIT

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Створіть метод Read([]byte) (int, error) для структурного типу MyReader.

func main() {
	reader.Validate(MyReader{})
}

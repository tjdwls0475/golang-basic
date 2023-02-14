package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
  // Here 'NewReader' function accepts 'string', and returns a 'pointer'
	r := strings.NewReader("I am Harry, and I am seeing Heyon.")
  // 'make' is a built-in function in Golang that allows you to create slices, maps, and channels. The main purpose of it is to allocate a new, empty object and return a pointer.
	b := make([]byte, 4)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}

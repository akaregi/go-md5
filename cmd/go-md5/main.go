package main

import (
	"fmt"
)

func main() {
	text := "abcde"
	binary := []byte(text)

	binary = append(binary, 0x80)

	for len(binary)%512 != 448 {
		binary = append(binary, 0x00)
	}

	binary = append(binary, 0x00)

	fmt.Printf("binary: %d\n", binary)
	fmt.Printf("hex:    %d\n", binary)
	fmt.Printf("length: %d\n", len(binary))
}

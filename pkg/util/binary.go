package util

import "fmt"

func ToBinaryExpr(text string) (binary []string) {
	for _, c := range text {
		binary = append(binary, fmt.Sprintf("%b", c))
	}

	return
}

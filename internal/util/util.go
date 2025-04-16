package util

import "fmt"

// ReverseString returns the input string reversed rune-wise
func ReverseString(s string) string {
	fmt.Println("V1")
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

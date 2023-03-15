package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}

	fmt.Println(maxProduct(a))

}

func maxProduct(words []string) int {
	sort.Strings(words)
	result := 0
	for i, j := len(words)-1, len(words)-2; j >= 0; i, j = i-1, j-1 {
		
	}
	return result
}

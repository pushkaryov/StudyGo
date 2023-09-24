package main

import (
	"fmt"
	"sort"
)

/**
дан слайс состоящий из строк. Необходимо напечатать true тогда и только
тогда, когда все слова в слайсе отсортированы лексикографически относительно
друг друга.
*/

func main() {
	words := []string{"apple", "banana", "cherry"}
	result := isSorted(words)
	fmt.Println(result)
}

func isSorted(input []string) bool {
	sorted := make([]string, len(input))
	copy(sorted, input)
	sort.Strings(sorted)
	return equal(input, sorted)
}

func equal(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

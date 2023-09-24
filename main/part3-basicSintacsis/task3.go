package main

import "fmt"

/**
Нам дана строка следующего вида “съешь ещё этих мягких французских булок, да
выпей чаю”. Используя тип данных map посчитайте количество повторений
символов в этой строке. В результате выведите на экран список символ -
количество повторений
*/

func main() {
	stringa := "съешь ещё этих мягких французских булок, да выпей чаю"

	characterCount := countCharacters(stringa)

	fmt.Println("Символ - Количество повторений:")
	for char, count := range characterCount {
		fmt.Printf("%c - %d\n", char, count)
	}
}

func countCharacters(input string) map[rune]int {
	charCount := make(map[rune]int)

	for _, char := range input {
		charCount[char]++
	}

	return charCount
}

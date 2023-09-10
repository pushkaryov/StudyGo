package main

import "fmt"

func main() {

	fmt.Print("Enter a number: ")
	var f float64
	fmt.Scanf("%f", &f)

	output := f * 0.3048

	fmt.Println(output)
}

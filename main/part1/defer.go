package main

import "fmt"

func end() {
	fmt.Println("functions end")
}

func main() {

	defer end()

	num := 5

	defer func(x int) {
		fmt.Println(x)
	}(num)

	num = 20

	//5
	//functions end

	//Выполняется в обратном порядке
}

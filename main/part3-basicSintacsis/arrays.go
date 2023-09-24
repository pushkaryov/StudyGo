package main

import "fmt"

func main() {

	fmt.Println("Тут будет выведен идентификатор”")

	var x []int64 = []int64{213, 123, 12}
	x = append(x, 22, 33, 33, 55)

	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println(x)

}

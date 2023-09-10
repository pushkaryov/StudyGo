package main

import "fmt"

func main() {

	var m1 map[int32]bool
	var m2 map[string]string

	m3 := make(map[int]int)

	ages := map[string]int{
		"Andrey": 32,
		"Vasya":  33,
		///
	}
	ages["Anton"] = 3
	ages["Anton"] = 35555

	age := ages["Andrey"]

	fmt.Println(age,
		ages["Anton"],
		m1,
		m2,
		m3)
}

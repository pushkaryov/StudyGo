package main

import "fmt"

func main() {

	type Point struct {
		X int
		Y int
	}

	p := Point{
		X: 5,
		Y: 3,
	}

	p = Point{3, 5}

	fmt.Println(p.Y)
	fmt.Println(p.X)

	p.X = 54

	fmt.Println(p.X)
}

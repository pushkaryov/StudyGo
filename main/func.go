package main

import "fmt"

func main() {
	myString := addPrefix("my_string")
	fmt.Println(myString)

	stringWithError, err := addPrefixWithErr("stringWithError")
	fmt.Println(stringWithError)
	fmt.Println(err)

}

func addPrefix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithErr(origin string) (string, error) {
	return "prefix_" + origin, nil
}

func addPrefixWithLen(origin string) (res string, lenght int) {
	res = "prefix_" + origin
	lenght = len(res)
	return
	//return res, lenght
}

package main

import "fmt"

func main() {
	const url = "www.darkforce.com"
	const c int = 2
	var firstPtr *string = new(string)
	*firstPtr = "Yoda"
	fmt.Println(*firstPtr)

	secondName := "Anakin"
	fmt.Println(secondName)
	ptr := &secondName
	fmt.Println(ptr, *ptr)
	secondName = "Darth Vader"
	fmt.Println(ptr, *ptr)
	fmt.Println(url)
	fmt.Println(float32(c) + 4.2)
	fmt.Println(c + 4)
}

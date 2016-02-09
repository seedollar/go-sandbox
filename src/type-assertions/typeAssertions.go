package main

import "fmt"

func main() {

	var i interface{} = "Empty Interface"

	defer func() {
		fmt.Println("About To Panic")
	}()

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Will cause a panic here
	b := i.(bool)
	fmt.Println(b)
}




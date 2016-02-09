package main

import "fmt"

func main() {
	defer fmt.Println("deferred line")
	fmt.Println("Before defer")
}

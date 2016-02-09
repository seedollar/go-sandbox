package main

import "fmt"

func sayWhat() {
	for x:=0; x<10; x++ {
		fmt.Printf("%d", x)
	}
}

func sayWhen() {
	for y:=100; y<110; y++ {
		fmt.Println(y)
	}
}

func sayWhere() {
	for z:=1000; z<1010; z++ {
		fmt.Println(z)
	}
}


func main() {
	sayWhat()
//	go sayWhen()
	go sayWhere()
}









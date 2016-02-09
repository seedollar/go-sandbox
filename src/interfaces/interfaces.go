package main

import "fmt"

type Person interface {
	GetGender() string
}

type Male struct {
	gender string
}

func (male Male) GetGender() string {
	return male.gender
}

type Female struct {
	gender string
}

func (female Female) GetGender() string {
	return female.gender
}

func describeInteface(person Person) {
	fmt.Printf("\n(%v, %T)", person, person)
}

func main() {
	var person1, person2 Person
	guy := Male{"male"}
	chic := Female{"female"}

	person1 = guy
	person2 = chic

	fmt.Println(person1.GetGender())
	fmt.Println(person2.GetGender())
	// print interface tuple
	describeInteface(person1)
	describeInteface(person2)

}



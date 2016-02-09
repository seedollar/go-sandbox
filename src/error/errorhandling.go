package main

import (
	"fmt"
	"time"
)

// Define an error struct
type MyError struct {
	When time.Time
	What string
}

// Implement the Error interface
func (me *MyError) Error() string {
	return fmt.Sprintf("error at %v, %s", me.When, me.What)
}

func run() error {
	return &MyError{time.Now(), "Time Lapse Error"}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

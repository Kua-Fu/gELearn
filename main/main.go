package main

import "fmt"

type t struct {
}

// f1 f2
func f1() error {
	fmt.Println("f1")
	fmt.Println("ff")
	return nil
}

func main() {
	f1()
	fmt.Println("hello world!")
	fmt.Println("hello sunshine!")
}

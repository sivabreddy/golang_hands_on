package main

import (
	"fmt"
)

// Declaring a pointer - Referencing & , accessing value via pointer - Dereferencing *

func pointers() {
	// b is the pointer to a, because it stores the address of a
    a := 10
	b := &a
	fmt.Println("the value of a is:", a) // prints 10
	fmt.Println("The address of a is:", b) // memory address of a
	fmt.Println("the value of a is:", *b) // dereferencing the pointer , it prints 10
}

// output
// the value of a is: 10
// The address of a is: 0xc000010090
// the value of a is: 10



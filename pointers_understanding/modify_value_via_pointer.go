package main

import (
	"fmt"
)

// modifying a value via pointer
func modifyValueViaPointer() {
	// b is the pointer to a, because it stores the address of a
    a := 10
	fmt.Println("the value of a is:", a) // prints 10

	b := &a
	fmt.Println("The address of a is:", b) // memory address of a

	*b = 20 //change the value at the memory address - dereferencing the pointer
    fmt.Println("the value of a is:", *b) // dereferencing the pointer , it prints 20
}

// output
// the value of a is: 10
// The address of a is: 0xc000090020
// the value of a is: 20
package main

import (
	"fmt"
)

// pointer to functions
// Function takes a pointer to an int.
func changeValue(num *int) {
    *num = 50 // Modify original value of 'a'
}

func pointersToFunctions() {
    a := 10
    changeValue(&a)  // Pass address of 'a'

    fmt.Println("Value of a after function call:", a) // output will be 50
}
// pointers - variables that store the address of a variable
// & operator - gives the address of a variable
// advantages 
// 1. Avoid unncessary copying of variables
// Avoid unncessary copying of variables - for large and complex values this may take up too much memory space unnecessarily.
// With pointers, only one value is stored in the memory and the address is passed around. i.e no cpoy is created.
// 2. directly mutate values
// pass a pointer(address)insted of a value to a function.
// The function can then directly edit the underlying value - no need to return the value.


// All values in Go have so called "Null value" i.e int - 0, float 0.0, string -""
// for a pointer, the null value is nil, means absence of an address value - i.e no address is stored in the pointer/no value in memory.

package main

import "fmt"

func main() {
	age := 32 //regular variable

	var agePointer *int
	agePointer = &age

	// insted of 2 lines we can write in shorthand agePointer := &age

	fmt.Println("Age:",age)
	// dereferencing the pointer here. using * infront of pointer gives the values stored in the address
	fmt.Println("AgePointer",*agePointer)
	editAgeToAdultYears(agePointer)
}

func editAgeToAdultYears(age *int) {
	//dereference the address and get the value in the address
	// return *age - 18
	*age = *age - 18
	fmt.Println("Adult Years:",*age)
}
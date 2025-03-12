package main

import (
	"fmt"
)


// struct pointers
type person struct {
	name string
	age  int
}

func structPointers() {
	p := &person{name: "John", age: 30} //pointer to struct
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println((p).name,(p).age) 
	fmt.Println((*p).name,(*p).age) 
}
// output
// &{John 30}
// {John 30}
// John 30
// John 30
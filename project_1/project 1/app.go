// first program in golang
// package clause - every go program must start with package clause
// can have multiple packages in a program
// Go comes with a standard library - with multiple packages
// package name can be anything , but main is reserved for main package. Its entry point for the program.
// module is a collection of packages
// module is nothing but a program that is built from a collection of packages
// commands to run a module 
// go mod init example.com/hello - creates a go.mod file
// go mode tidy - cleans up the go.mod file
// go build - build a program
// ./hello - run a program
package main
//fmt is package for formatting input and output
import "fmt"

func main () {
	fmt.Print("Hello World")
}
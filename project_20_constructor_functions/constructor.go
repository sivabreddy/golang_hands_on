package main

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// we can define struct inside or ourside the function
// Use first character uppercase to make it public to access in other packages
type User struct {
	firstName string
	lastName  string
	birthdate string
	//Time is also struct provided by the time package
	createdAt time.Time
}
// function attached to a struct is called method
// receiver is the variable that is used to call the method
func (u *User) outputUserDetails() { 
    fmt.Println(u.firstName, u.lastName, u.birthdate,u.createdAt)
}
// replace the first and last name with empty string in the struct
func (u *User) clearUserName() {
   u.firstName = ""
   u.lastName = ""
}

// creations or constructor function , returns a pointer
func newUser(userFirstName, userLastName, userBirthdate string) (*User, error){
	//validations
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, 
		errors.New("first name, last name and birthdate are required")
	}
	return &User {
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(), // comma is needed even if it is a last member
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

    // Here the "appUser" is a variable of type User
	var appUser *User
	
    appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}


    // Here we are calling the method on the appUser variable
	appUser.outputUserDetails() //output details
	appUser.clearUserName() // empty firstName and lastName
	appUser.outputUserDetails() //output details again

}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
// output
// @sivas-MacBook-Pro project_19_mutation_methods % ./constructor 
// Please enter your first name: siva
// Please enter your last name: bojja
// Please enter your birthdate (MM/DD/YYYY): 11/01/1988
// siva bojja 11/01/1988 2025-03-11 17:11:47.393671 +0530 IST m=+9.166852209
//   11/01/1988 2025-03-11 17:11:47.393671 +0530 IST m=+9.166852209
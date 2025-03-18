package main

import (
	"fmt"
	"time"
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
// To mutate the receiver, define the method with a pointer receiver (*T). 
// Without a pointer, the method operates on a copy of the receiver, leaving the original unchanged.

// mutation of the struct is possible
// copy of values are passed to the function
func (u *User) clearUserName() {
   u.firstName = ""
   u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

    // Here the "appUser" is a variable of type User
	var appUser User
    // order of the fields is important in shorthand notation
	// appUser := User { userFirstName, userLastName, userBirthdate, time.Now() }
	// In below long hand notation approach, if we don't have any values, we can skip those.
	appUser = User {
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(), // comma is needed even if it is a last memeber
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
// @sivas-MacBook-Pro project_19_mutation_methods % ./mutations 
// Please enter your first name: siva
// Please enter your last name: bojja
// Please enter your birthdate (MM/DD/YYYY): 11/01/1988
// siva bojja 11/01/1988 2025-03-11 17:11:47.393671 +0530 IST m=+9.166852209
//   11/01/1988 2025-03-11 17:11:47.393671 +0530 IST m=+9.166852209
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

	outputUserDetails(&appUser)

}

func outputUserDetails(u *User){
    fmt.Println(u.firstName, u.lastName, u.birthdate,u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
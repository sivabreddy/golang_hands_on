package user

import (
	"fmt"
	"time"
	"errors"
)

// we can define struct inside or ourside the function
// Use first character uppercase to make it public to access in other packages
type User struct {
	FirstName string
	LastName  string
	Birthdate string
	//Time is also struct provided by the time package
	CreatedAt time.Time
}
// function attached to a struct is called method
// receiver is the variable that is used to call the method
func (u *User) OutputUserDetails() { 
    fmt.Println(u.FirstName, u.LastName, u.Birthdate, u.CreatedAt)
}
// replace the first and last name with empty string in the struct
func (u *User) ClearUserName() {
   u.FirstName = ""
   u.LastName = ""
}


// creations or constructor function , returns a pointer
// try to give constructor function name as New <= it is a practice in production code
func New(userFirstName, userLastName, userBirthdate string) (*User, error){
	//validations
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, 
		errors.New("first name, last name and birthdate are required")
	}
	return &User {
		FirstName: userFirstName,
		LastName: userLastName,
		Birthdate: userBirthdate,
		CreatedAt: time.Now(), // comma is needed even if it is a last member
	}, nil
}
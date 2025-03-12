package main
import (
	"fmt"
	"example.com/constructor/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

    // Here the "appUser" is a variable of type User
	var appUser *user.User
	
    appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}


    // Here we are calling the method on the appUser variable
	appUser.OutputUserDetails() //output details
	appUser.ClearUserName() // empty firstName and lastName
	appUser.OutputUserDetails() //output details again

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
// output
// @sivas-MacBook-Pro project_19_mutation_methods % ./constructor 
// Please enter your first name: siva
// Please enter your last name: bojja
// Please enter your birthdate (MM/DD/YYYY): 11/01/1988
// siva bojja 11/01/1988 2025-03-11 17:11:47.393671 +0530 IST m=+9.166852209
//   11/01/1988 2025-03-11 17:11:47.393671 +0530 IST m=+9.166852209
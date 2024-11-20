package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	secondName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user

	var appUser = user{
		firstName: firstName,
		lastName:  secondName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	// fmt.Println(firstName, lastName, birthdate)
	printIt(appUser)

}

func printIt(user user) {
	fmt.Print("Hello", user.firstName)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

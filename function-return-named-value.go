package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Devri"
	middleName = "Wirlandika"
	lastName = "SI20B"
	return
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Printf(firstName)
	fmt.Printf(middleName)
	fmt.Printf(lastName)
}
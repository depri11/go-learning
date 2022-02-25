package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) :",numberA)
	fmt.Println("numberA (value) :",&numberA)
	fmt.Println("numberB (value) :",*numberB)
	fmt.Println("numberB (value) :",&numberB)

	numberA = 5

fmt.Println("numberA (value)   :", numberA)
fmt.Println("numberA (address) :", &numberA)
fmt.Println("numberB (value)   :", *numberB)
fmt.Println("numberB (address) :", numberB)

var alamat = Address {
	City: "Subang",
	Province: "Jawa Barat",
	Country: "",
}

ChangeCountryToIndonesia(&alamat)
fmt.Println(alamat)
	
}
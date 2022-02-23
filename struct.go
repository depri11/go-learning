package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
	
}

func (a Customer) abcDE(){
	fmt.Println("Hai", a.Name)
}

func main() {
	var devri Customer
	devri.Name = "Devri"
	devri.Address = "Jakarta"
	devri.Age = 21

	devri.sayHi("Devri")
	devri.abcDE()

	// fmt.Println(dev)
}
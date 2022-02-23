package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(name HasName) {
	fmt.Println("Hello", name.GetName())
}

func main() {
	var dev Person
	dev.Name = "Devri"
	SayHello(dev)

	cat := Animal {
		Name: "Push",
	}
	SayHello(cat)
}
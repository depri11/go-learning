package main

import "fmt"

func main() {
	name := "devri"
	counter := 0

	increment := func() {
		name := "wirlandika"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
		
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
package main

import "fmt"

func main() {
	var fruits = []string{"apple", "orange", "banana", "melon"}
	fmt.Println(fruits[3])

	var newFruits = fruits[0:2]
	fmt.Println(newFruits)

	// dst := make([]string, 3)
	// src := []string{"apple", "orange", "banana", "watermelon"}
	// n := copy(dst, src)

	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)

	dst := []string{"apple", "orange"}
	src := []string{"nasi", "sayur", "lemon", "pepaya"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

}
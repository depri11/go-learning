package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Error dengan message:", message)
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	
}

func main() {
	runApp(false)
}
package main

import "fmt"

func main() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	
	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	//map and slice
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "Male"},
		map[string]string{"name": "chicken red", "gender": "Male"},
		map[string]string{"name": "chicken yellow", "gender": "Female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"],chicken["name"])
	}

}
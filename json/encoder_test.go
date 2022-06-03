package json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customerOut.json")

	customer := Customer{
		FirstName: "Devri",
		LastName:  "Wirlandika",
		Age:       21,
	}

	err := json.NewEncoder(writer).Encode(customer)
	if err != nil {
		panic(err)
	}
}

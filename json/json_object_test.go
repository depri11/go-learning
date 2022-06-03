package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	LastName  string
	Age       int
	Married   bool
	Hobbies   []string
	Addresses []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Dev",
		LastName: "Wirlandika",
		Age:      21,
		Married:  false,
	}

	byte, _ := json.Marshal(customer)

	fmt.Println(string(byte))
}

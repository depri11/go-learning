package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONA(t *testing.T) {
	customer := Customer{
		FirstName: "Dev",
		LastName:  "Wirlandika",
		Age:       21,
		Married:   false,
		Hobbies:   []string{"sport", "music", "Coding"},
	}

	byte, _ := json.Marshal(customer)

	fmt.Println(string(byte))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FistName":"Dev","LastName":"Wirlandika","Age":21,"Married":false,"Hobbies":["sport","music","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}

	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Dev",
		Addresses: []Address{
			{
				Street:     "Karawang",
				Country:    "Indonesia",
				PostalCode: "41731",
			},
			{
				Street:     "Jakarta",
				Country:    "Indonesia",
				PostalCode: "14794",
			},
		},
	}

	byte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Dev","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Karawang","Country":"Indonesia","PostalCode":"41731"},{"Street":"Jakarta","Country":"Indonesia","PostalCode":"14794"}]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}

	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Karawang","Country":"Indonesia","PostalCode":"41731"},{"Street":"Jakarta","Country":"Indonesia","PostalCode":"14794"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {

	address := []Address{
		{
			Street:     "Karawang",
			Country:    "Indonesia",
			PostalCode: "41731",
		},
		{
			Street:     "Jakarta",
			Country:    "Indonesia",
			PostalCode: "14794",
		},
	}

	byte, err := json.Marshal(address)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

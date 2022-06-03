package json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")

	customer := &Customer{}

	err := json.NewDecoder(reader).Decode(&customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

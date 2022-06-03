package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id": "1", "name": "Sepatu", "price": "15000"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncoding(t *testing.T) {
	product := map[string]interface{}{
		"id":    "1",
		"name":  "Sepatu",
		"price": "15000",
	}

	byte, _ := json.Marshal(&product)
	fmt.Println(string(byte))
}

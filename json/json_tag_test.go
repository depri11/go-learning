package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "1",
		Name:     "Sepatu",
		ImageURL: "http://image.com/image.jpg",
	}

	byte, _ := json.Marshal(product)
	fmt.Println(string(byte))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"1","name":"Sepatu","image_url":"http://image.com/image.jpg"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}

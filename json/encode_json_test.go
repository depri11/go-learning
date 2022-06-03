package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(string(byte)))
}

func TestEncode(t *testing.T) {
	logJson("Dev")
	logJson(1)
	logJson(true)
	logJson([]string{"Dev", "Wirlandika"})
}

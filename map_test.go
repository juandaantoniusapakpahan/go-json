package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeMapJson(t *testing.T) {
	type M map[string]interface{}
	dataMap := M{
		"id":    "001",
		"name":  "Mac Book Pro 2020",
		"price": 2000000,
	}

	result, err := json.Marshal(dataMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}
func TestDecodeMapJson(t *testing.T) {
	jsonData := `{"id":"001","name":"Mac Book Pro 2020","price":2000000}`

	var M map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &M)
	if err != nil {
		panic(err)
	}

	fmt.Println(M)
}

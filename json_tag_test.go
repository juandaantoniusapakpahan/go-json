package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`        // json tag
	Name     string `json:"name"`      // json tag
	ImageUrl string `json:"image_url"` // json tag
}

func TestEncodeJSONTag(t *testing.T) {
	product := Product{
		Id:       "0001",
		Name:     "Mac Book Pro 2020",
		ImageUrl: "http://localhost:8080/image.png",
	}

	result, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}

func TestDecodeJSONTag(t *testing.T) {
	jsonData := `{"id":"0001","name":"Mac Book Pro 2020","image_url":"http://localhost:8080/image.png"}`
	product := &Product{}

	err := json.Unmarshal([]byte(jsonData), product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}

package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func EncodeJson(value interface{}) {
	result, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}

func TestEncodeJson(t *testing.T) {
	EncodeJson("test")
	EncodeJson(1)
	EncodeJson(false)
	EncodeJson(true)
	EncodeJson([]string{"Lontong", "Tahu isi", "Robert"})

}

package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingDecoder(t *testing.T) {
	reader, err := os.Open("customer.json")
	if err != nil {
		panic(err)
	}

	cstm := &Customer{}
	decoder := json.NewDecoder(reader)
	decoder.Decode(cstm)

	fmt.Println(cstm)
}

package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJsonObject(t *testing.T) {
	constumer := Customer{
		FirstName:  "Richard",
		MiddleName: "Param",
		LastName:   "Smith",
	}

	result, err := json.Marshal(constumer)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}

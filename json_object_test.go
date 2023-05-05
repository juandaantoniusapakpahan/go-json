package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Address    string
	Country    string
	PostalCode int
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
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

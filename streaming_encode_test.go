package gojson

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestEncoderStreming(t *testing.T) {
	file, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(file)

	ctmr := Customer{
		FirstName:  "Lola",
		MiddleName: "Serly",
		LastName:   "Nidle",
		Age:        43,
		Married:    true,
		Hobbies:    []string{"Makan", "Minum"},
		Addresses: []Address{
			{
				Address:    "New York",
				Country:    "USA",
				PostalCode: 23232,
			}, {
				Address:    "London",
				Country:    "England",
				PostalCode: 89595,
			},
			{
				Address:    "Bangkok",
				Country:    "Thailand",
				PostalCode: 239500,
			},
			{
				Address:    "Bali",
				Country:    "Indonesia",
				PostalCode: 80808,
			}},
	}

	err := encoder.Encode(ctmr)
	if err != nil {
		panic(err)
	}

	result, _ := os.Open("CustomerOut.json")
	data := make([]byte, 100)
	count, err := result.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

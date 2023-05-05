package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeJsonArrayHobbies(t *testing.T) {
	ctmr := Customer{
		FirstName:  "Lola",
		MiddleName: "Serly",
		LastName:   "Nidle",
		Age:        43,
		Married:    true,
		Hobbies:    []string{"Makan", "Minum"},
	}

	result, err := json.Marshal(ctmr)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}
func TestDecodeJsonArrayHobbies(t *testing.T) {
	dataJson := `{"FirstName":"Lola","MiddleName":"Serly","LastName":"Nidle","Age":43,"Married":true,"Hobbies":["Makan","Minum"],"Addresses":null}`

	ctmr := &Customer{}
	err := json.Unmarshal([]byte(dataJson), ctmr)
	if err != nil {
		panic(err)
	}

	fmt.Println(ctmr)
}

func TestEncodeJsonArrayAddress(t *testing.T) {
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

	result, err := json.Marshal(ctmr)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
	fmt.Println(string(result))

}

func TestDecodeJsonArrayAddress(t *testing.T) {
	dataJson := `{"FirstName":"Lola","MiddleName":"Serly","LastName":"Nidle","Age":43,"Married":true,"Hobbies":["Makan","Minum"],"Addresses":[{"Address":"New York","Country":"USA","PostalCode":23232},{"Address":"London","Country":"England","PostalCode":89595},{"Address":"Bangkok","Country":"Thailand","PostalCode":239500},{"Address":"Bali","Country":"Indonesia","PostalCode":80808}]}`
	addressJson := `[{"Address":"New York","Country":"USA","PostalCode":23232},{"Address":"London","Country":"England","PostalCode":89595},{"Address":"Bangkok","Country":"Thailand","PostalCode":239500},{"Address":"Bali","Country":"Indonesia","PostalCode":80808}]`
	ctmr := &Customer{}
	adrs := &[]Address{}
	err := json.Unmarshal([]byte(dataJson), ctmr)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(addressJson), adrs)
	if err != nil {
		panic(err)
	}

	//fmt.Println(ctmr)
	fmt.Println(adrs)
}

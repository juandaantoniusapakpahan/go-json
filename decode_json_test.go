package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	dataJson := `{"FirstName":"Richardoe", "MiddleName":"Jolo", "LastName":"Gianto", "Age":23, "Married":true}`
	ctm := &Customer{}
	datajsonByte := []byte(dataJson)
	err := json.Unmarshal(datajsonByte, ctm)

	if err != nil {
		panic(err)
	}
	fmt.Println(ctm)
	fmt.Println(ctm.FirstName)
	fmt.Println(ctm.Age)
	fmt.Println(ctm.Married)

}

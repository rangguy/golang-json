package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Rangga",
		MiddleName: "Dwi",
		LastName:   "Mahedra",
		Age:        18,
		Married:    false,
		Hobbies:    []string{"Game", "Coding", "Running"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Rangga","MiddleName":"Dwi","LastName":"Mahedra","Age":18,"Married":false,"Hobbies":["Game","Coding","Running"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Rangga",
		Addresses: []Address{
			{
				Street:     "Street",
				Country:    "US",
				PostalCode: "999",
			},
			{
				Street:     "Jalan",
				Country:    "Indonesia",
				PostalCode: "123",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Rangga","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Street","Country":"US","PostalCode":"999"},{"Street":"Jalan","Country":"Indonesia","PostalCode":"123"}]}
`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArray(t *testing.T) {
	jsonString := `[{"Street":"Street","Country":"US","PostalCode":"999"},{"Street":"Jalan","Country":"Indonesia","PostalCode":"123"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)

}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Street",
			Country:    "US",
			PostalCode: "999",
		},
		{
			Street:     "Jalan",
			Country:    "Indonesia",
			PostalCode: "123",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}

package main

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

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Rangga",
		MiddleName: "Dwi",
		LastName:   "Mahedra",
		Age:        18,
		Married:    false,
	}

	bytes, err := json.Marshal(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

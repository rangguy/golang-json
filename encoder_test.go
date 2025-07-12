package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		FirstName:  "Mahendra",
		MiddleName: "Dwi",
		LastName:   "Rangga",
	}

	_ = encoder.Encode(customer)

	fmt.Println(customer)
}

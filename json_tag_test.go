package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		ID:       "1",
		Name:     "Rumah",
		ImageURL: "https://www.google.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}``

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"1","name":"Rumah","image_url":"https://www.google.com"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}

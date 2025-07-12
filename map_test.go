package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"1","name":"Rumah","price":"20000"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "1",
		"name":  "Rumah",
		"price": "20000",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

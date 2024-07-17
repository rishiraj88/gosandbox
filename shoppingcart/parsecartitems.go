// to implement the getCartFromJson() method.
package main

import (
	"encoding/json"
	"fmt"
)

const showExpectedResult = false;
const showHints = false;

type cartItem struct {
	Name  string "json:name"
	Price float64 "json:price"
	Quantity int "json:quantity"
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
    var cart []cartItem

	err := json.Unmarshal([]byte(jsonString),&cart)
	if nil != err {
		panic(err)
	}
	return cart
}

func main() {
	jsonString := 
	`[{"name":"apple","price":2.99,"quantity":3},
  	{"name":"orange","price":1.50,"quantity":8},
  	{"name":"banana","price":0.49,"quantity":12}]`
	
	fmt.Println(getCartFromJson(jsonString))
}
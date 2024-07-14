package main

import "fmt"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// Converts a slice of strings to a map object.
func convertToMap1(items []string) map[string]float64 {

    // Create a map object
	result := make(map[string]float64)
    
    price := 100.0 / float64(len(items))
    for k := range items {
     result[items[k]] = price
    }
    return result
}

// Converts a slice of strings to a map object.
func convertToMap2(items []string) map[string]float64 {

    // Create a map object
	result := make(map[string]float64)
    
	element := 100 / float64(len(items))
	for _,fruit := range items {
		result[fruit] = element
	}

	return result
}

func main(){
	slice := []string{"apple", "banana", "orange", "date"}
	fmt.Println(convertToMap2(slice))
}
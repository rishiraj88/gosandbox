package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {

	// Convert the first string to a float64
	val1, err1 := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err1 != nil {
		panic(err1)
	}
	// Convert the second string to a float64
	val2, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err2 != nil {
		panic(err2)
	}
	// Calculate and return the result
	sum := math.Round((val1 + val2)*100) / 100
	fmt.Println("The sum is:", sum)
	return sum
}

func main(){
	calculate("23.3","32.3")
}
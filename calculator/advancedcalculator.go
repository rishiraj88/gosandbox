// Write your answer here, and then test your code.

package main

import (
	"fmt"
	"strconv"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

func main(){
	value1 := "10"
	value2 := "5.5"
	operation := "+"
	fmt.Println(calculate(value1, value2, operation))
}
// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	val1 := convertInputToValue(input1)
	val2 := convertInputToValue(input2)
	
	switch operation {
		case "+": return addValues(val1,val2)
		case "-": return subtractValues(val1,val2)
		case "*": return multiplyValues(val1,val2)
		case "/": return divideValues(val1,val2)
		default: return -99-99990909
	}
}

func convertInputToValue(input string) float64 {
	val,err := strconv.ParseFloat(input,64)
	if(nil != err) {
		panic(err)
	}
	return val
}

func addValues(value1, value2 float64) float64 {

	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}


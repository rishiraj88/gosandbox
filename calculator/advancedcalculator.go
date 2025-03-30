package main

import (
	"fmt"
	"strconv"
)

func main(){
	value1 := "10"
	value2 := "4.5"
	operation := "*"
	result,status :=	calculate(value1, value2, operation)
	fmt.Println(result)
	if "" != status {
		panic(status)
	} 
}
// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) (float64,string) {
	val1 := convertInputToValue(input1)
	val2 := convertInputToValue(input2)
	
	switch operation {
		case "+": return add(val1,val2),""
		case "-": return subtract(val1,val2),""
		case "*": return multiply(val1,val2),""
		case "/": return divide(val1,val2),""
		default: return -99-99990909,"invalid arguments"
	}
}

func convertInputToValue(input string) float64 {
	val,err := strconv.ParseFloat(input,64)
	var errMsg string
	if(nil != err) {
		errMsg = fmt.Sprintf("Error while parsing an input. %v must be a number.",input)
		panic(errMsg)
	}
	return val
}

func add(value1, value2 float64) float64 {

	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}


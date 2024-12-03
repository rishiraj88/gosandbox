package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text: ")
	input, _ := reader.ReadString(byte('\n'))
	fmt.Println("You have entered:", input)
	input, _ = reader.ReadString(byte('\n'))
}

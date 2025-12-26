package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){

	fmt.Println("Entre com a operação (use este formato: 2+2):")
	var input string
	// entrada de valores no terminal
	fmt.Scan(&input)

	operation := strings.Split(input, "")

	result := getResult(operation)

	fmt.Printf("%s %s %s = %d\n", operation[0], operation[1], operation[2], result)
}

func getResult(operation []string ) int {
	num1, _ := strconv.Atoi(operation[0])  
	num2, _ := strconv.Atoi(operation[2])

	switch operation[1] {
	case "+":
		return num1 + num2 
	case "-":
		return num1 - num2 
	case "*":
		return num1 * num2 
	case "/":
		return num1 / num2
	default:
		panic("Operator not valid") 

	}
}

// as 4 operações
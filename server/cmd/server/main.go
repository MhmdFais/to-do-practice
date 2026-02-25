package main

import "fmt"

func main() {
	secondaryFunction()
	fmt.Println()
	var result int = addTwoNumbers(5, 10)  // result := addTwoNumbers(5, 10)
	fmt.Printf("The result of adding 5 and 10 is: %d\n", result)
	fmt.Println()
	age := 20
	fmt.Println(ifStatementExample(age))
}

func secondaryFunction() {
	fmt.Println("This is a secondary function.")
}

func addTwoNumbers(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

func ifStatementExample(age int) string {
	if age >= 18 {
		return "You are an adult."
	} else {
		return "You are a minor."
	}
}
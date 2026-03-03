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
	fmt.Println()
	arrayExample()
	fmt.Println()
	sliceExample()
	fmt.Println()
	forLoopExample()
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

func arrayExample(){
	var numbers[5]int

	// numbes := [5]int{10, 20, 30, 40, 50} // Array literal

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Array of integers:", numbers)
	fmt.Println(numbers[0:5])
}

func sliceExample() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice of integers:", slice)
	
	slice = append(slice, 60)
	fmt.Println("Slice after appending 60:", slice)
}

func forLoopExample() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
}
package main

import (
	"errors"
	"fmt"
)

func main() {

	scores := []int{10, 30, 5, 15, 20}
	total := sum(scores)
	fmt.Println(total)

	result, err := calculate(10, 2, "p")
	// cek apakah ada error
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total = total + number
	}
	return total
}

func calculate(numberOne int, numberTwo int, operation string) (int, error) {

	var result int
	var error error

	switch operation {
	case "+":
		result = numberOne + numberTwo
	case "-":
		result = numberOne - numberTwo
	case "*":
		result = numberOne * numberTwo
	case "/":
		result = numberOne / numberTwo
	default:
		error = errors.New("unknown operation")
	}
	return result, error
}

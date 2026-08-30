package main

import "fmt"

func double(value int) int {
	return value * 2
}

func divide(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", dividend)
	}
	return dividend / divisor, nil
}

func main() {
	fmt.Println(double(21))

	result, err := divide(20, 2)
	fmt.Printf("result = %v error = %v\n", result, err)

	result, err = divide(20, 0)
	fmt.Printf("result = %v error = %v\n", result, err)

	quotient, _ := divide(21, 3)
	fmt.Println("\nIgnoring the error. Quotient:", quotient)

	if remainder := 10 % 3; remainder != 0 {
		fmt.Println("\n10 is not divisible by 3, remainder:", remainder)
	}

	count := 0
	for count < 3 {
		count++
	}
	fmt.Println("\ncounted to:", count)

	switch count {
	case 3:
		fmt.Println("count is 3, breaking switch statement")
	default:
		fmt.Println("count is something else")
	}
}

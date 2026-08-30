package main

import (
	"errors"
	"fmt"
)

var errDivideByZero = errors.New("divide by zero")

func divide(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errDivideByZero
	}
	return dividend / divisor, nil
}

func describe(top int, bottom int) error {
	defer fmt.Println("describe done")

	result, err := divide(top, bottom)
	if err != nil {
		return fmt.Errorf("describe %d/%d: %w", top, bottom, err)
	}
	fmt.Println("result:", result)
	return nil
}

func main() {
	err := describe(10, 2)
	fmt.Println(err)
	fmt.Println("was divided by zero?", errors.Is(err, errDivideByZero))

	fmt.Println()

	err = describe(10, 0)
	fmt.Println(err)
	fmt.Println("was divided by zero?", errors.Is(err, errDivideByZero))
}

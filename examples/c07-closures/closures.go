package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func countClosure() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	double := func(number int) int { return number * 2 }
	fmt.Println("double(5):", double(5))

	fmt.Println("sum(1, 2, 3):", sum(1, 2, 3))

	values := []int{4, 5, 6}
	fmt.Println("from slice []int{4, 5, 6} sum(values...):", sum(values...))

	counter := countClosure()
	counter()
	counter()
	fmt.Println("After three calls:", counter())
}

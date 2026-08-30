package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)
	fmt.Printf("slice: %v length = %d\n", numbers, len(numbers))

	ages := map[string]int{"alice": 30}
	ages["bob"] = 25
	fmt.Printf("map: %v\n", ages)
	fmt.Printf("missing key: %v\n", ages["nobody"])
	age, found := ages["nobody"]
	fmt.Printf("age = %v found = %v\n", age, found)

	for i, number := range numbers {
		fmt.Printf("numbers[%d] = %d\n", i, number)
	}
	for name, age := range ages {
		fmt.Printf("ages[%q] = %d\n", name, age)
	}

	aliased := numbers
	aliased[0] = 99
	fmt.Printf("numbers = %v aliased = %v\n", numbers, aliased)
}

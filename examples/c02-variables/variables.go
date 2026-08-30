package main

import "fmt"

var canDefineOutsideFunctions = true
// cantDefineOutsideFunctions := false
// ^ This will cause an error: expected declaration, found cantDefineOutsideFunctions

func main() {
	fmt.Println("-- Zero values --")
	var i int
	var f float64
	var str string
	var b bool

	fmt.Printf("int: %v (%T)\n", i, i)
	fmt.Printf("float64: %v (%T)\n", f, f)
	fmt.Printf("string: %q (%T)\n", str, str)
	fmt.Printf("bool: %v (%T)\n", b, b)

	fmt.Println("-- declarations --")
	fmt.Printf("canDefineOutsideFunctions: %v (%T)\n", canDefineOutsideFunctions, canDefineOutsideFunctions)

	fmt.Println("-- type strictness --")
	x := 42
	y := 4.2

	// sum := x + y
	// ^ invalid operation: x + y (mismatched types int and float64)

	sum := float64(x) + y
	fmt.Printf("sum: %v (%T)\n", sum, sum)

	fmt.Println("-- constants --")
	const answer = 42
	fmt.Printf("answer * 1.5: %f (%T)\n", answer*1.5, answer*1.5)
}
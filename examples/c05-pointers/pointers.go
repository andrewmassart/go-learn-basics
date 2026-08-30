package main

import "fmt"

func main() {
	count := 10
	pointer := &count

	fmt.Println("memory address:", pointer)
	fmt.Println("value at address", *pointer)

	*pointer = 20
	fmt.Println("count after write through pointer:", count)

	var unset *int
	fmt.Println("unset pointer is nil:", unset == nil)
}

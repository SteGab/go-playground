package main

import (
	"fmt"
	"sg/testapp/calc"
)

func main() {
	fmt.Println("Hello, World!")
	resAdd := calc.Add(1, 2)
	resSub := calc.Subtract(1, 2)
	resMul := calc.Multiply(1, 2)
	resDiv := calc.Divide(1, 2)
	fmt.Println("Addition: ", resAdd)
	fmt.Println("Subtraction: ", resSub)
	fmt.Println("Multiplication: ", resMul)
	fmt.Println("Division: ", resDiv)

}

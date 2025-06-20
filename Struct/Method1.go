package main

import "fmt"

type Calculator struct {
	Num1 int
	Num2 int
}

func (c Calculator) Multiply() int {
	return c.Num1 * c.Num2
}


func main() {
	mul := Calculator{24, 35}
	fmt.Printf("%v", mul.Multiply())
} 

package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var numerator int = 11
	var denominator int = 11
	var div, mod, err = ArithmeticOperation(numerator, denominator)

	// Handle error first
	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	// Check the value of mod
	switch {
	case mod == 0:
		fmt.Printf("The division was exact. Division: %v\n", div)
	default:
		fmt.Printf("Division: %v, Modulus: %v\n", div, mod)
	}

	// Additional switch for mod
	switch mod {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}
}

func ArithmeticOperation(numerator int, denominator int) (int, int, error) {
	t := time.Now()// Record the start time of the operation
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}
	var div int = numerator / denominator
	var mod int = numerator % denominator
	fmt.Println("Elapsed time:", time.Since(t))// Print the time taken to perform the operations
	return div, mod, err
	
}
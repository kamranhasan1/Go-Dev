package main

import (
    "errors"
    "fmt"
)

func main() {
    var numerator = 12
    var denominator = 5
    var div, mod, err = ArthOperation(numerator, denominator)  

    if err != nil {
        fmt.Println(err) // Prints the error message
    }else if mod == 0 {
		fmt.Printf("the result of ind div is %v",mod)
		
	} else {
        fmt.Printf("Division: %v, Modulus: %v\n", div, mod) 
    }
}

func ArthOperation(numerator int, denominator int) (int, int, error) {
	var err error
    if denominator == 0 {
    	err = errors.New("denominator cannot be zero") 
		return 0, 0,err
    }
    var div int = numerator / denominator
    var mod int = numerator % denominator

    return div, mod, err
}

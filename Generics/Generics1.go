package main

import (
	"fmt"
)

type Number interface {
    int64 | float64
}




// SumIntsOrFloats is a generic function to sum values in a map
func SumIntsOrFloats[K comparable, V int64 | float64 ](m map[K] V ) V {
	var s V
	for _, v := range m { 
		s += v
	}
	return s
}
 

func main() {

	// Initialize a map for the integer values
	ints := map[string]int64{

		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}


	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",

    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))


	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))


 
}

package main

import "fmt"

func main() {
 	var arr1 = [5]float64{2, 3, 4, 5, 6}
	
 	fmt.Printf("the memory location in arr1 %p", &arr1)
	
	// Calling Sqr function with address of arr1 and storing the result
	var res [5]float64 = Sqr(&arr1)
	
 	fmt.Printf("\n%v", res)
	
	// Printing the original arr1 to show it was modified
	fmt.Printf("\nthe value of arr1 is %v", arr1)
}

/*Sqr function takes a pointer to an array of 5 float64 elements
 and returns an array of 5 float64 elements */
func Sqr(arr2 *[5]float64) [5]float64 {
 	fmt.Printf("\nthe memory location in arr2 is %p", &arr2)
	
 	for i := range arr2 {
 		arr2[i] = arr2[i] * arr2[i]
	}
	
 	return *arr2
}
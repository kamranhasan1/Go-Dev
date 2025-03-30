package main

import (
	"fmt"
)

func main() {
	str := []rune("kamran")// Converts string to a slice of runes (Unicode code points)
	ind := (str[0])
	fmt.Printf("%v\n",ind)
	fmt.Println("Using []rune:")

	for i, v := range str {
		fmt.Printf("Index: %d, Rune: %c, Value: %d\n", i, v, v)
	}

	strBytes := []byte("kamran")// Converts string to a slice of bytes (UTF-8 encoding)
	fmt.Println("\nUsing []byte:")
	for i, v := range strBytes {
		fmt.Printf("Index: %d, Byte: %c, Value: %d\n", i, v, v)
	}

	// we can declare our rune type like this also
	var myrune = 'd'
	fmt.Printf("%v",myrune)
}

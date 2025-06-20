 package main

import (
	"fmt"
	"time"
)

// func task(id int){
// 	 fmt.Println("doing task no:- ",id)
// }

func main() {
	for i := 0; i < 10; i++ {
	// go task(i)

		// Using an anonymous function
		go func(i int) { 
			fmt.Println("task no",i)
		}(i)
	}
	time.Sleep(time.Millisecond * 10)// To prevent the program from exiting while goroutines are still running this is not right solution 
}

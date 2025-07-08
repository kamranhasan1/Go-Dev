package main

import (
	"fmt"
	// "math/rand"
	// "time" 
)

func main() {

	messageChan := make(chan string)
	go Printmessage(messageChan)
	fmt.Println(<-messageChan) // Waits and prints "Ping"

	// numChan := make(chan int)
	// go processNum(numChan)
	// for num := range numChan { // Processes numbers until channel closes
	// 	fmt.Println("Processing number", num) 
	// }


	result := make(chan int)
	go sum(result,4,5)

	res:= <-result
	fmt.Println(res) 



}





func Printmessage(c chan string) {
	defer close(c)   
	c <- "Ping"     
}





// func processNum(numChan chan int) {
// 	defer close(numChan)  
// 	for {
// 		numChan <- rand.Intn(100) // Sends random numbers infinitely
// 		time.Sleep(time.Second)    
// 	}
// }



func sum(result chan int,num1,num2 int)  {
	defer close(result)

	resultSum:= num1+num2 

	result <- resultSum
	
}
/* Channels are a powerful feature in Go that enable communication and synchronization between goroutines (lightweight threads).
 They provide a way for goroutines to send and receive values with each other.*/



 

 package main

 import (
	 "fmt"
	 "time"
 )
 
 func main() {
	 // Create a buffered channel with capacity 4 (can hold 4 ints without blocking)
	 var c = make(chan int, 4)
	 
 	 go process(c)
	 
	 // Range over the channel to receive values until it's closed
 	 for i := range c {
		 fmt.Println(i)  
	 }
	 
 	 time.Sleep(1 * time.Second)
 }
 
 func process(c chan int) {
	  
	 defer fmt.Println("Channel processing complete") 
	 defer close(c)                                  
 
	 // Send values  to the channel
	 for i := 0; i < 5; i++ {
		 c <- i 
	 }
  }
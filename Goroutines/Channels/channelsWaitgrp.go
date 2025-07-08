package main

import "fmt"

func main() {
	done := make(chan bool)   
	go task(done)            
	<-done                  // Blocks until 'done' receives a value
}

//gorutine syncronizer 
func task(done chan bool) {
	defer close(done)
	
	defer func() { done <- true }() // Ensures 'true' is sent before exiting
	fmt.Println("Processing...")   
	 
}
// with Concurrency 
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}//they are like counter
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)// to call the concurrency we'll have to add go with the func call
	}
	
	wg.Wait()//wait for the counter to go back to the zero meaning NULL
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
}

func dbCall(i int) {
	
	// simulate DB call delay
	delay := rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	
	fmt.Printf("The result from the database is: %s\n", dbData[i])

	defer wg.Done()// this decrements the counter 
}

  

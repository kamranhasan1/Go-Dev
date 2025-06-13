 /* The primary purpose of mutexes is to prevent race conditions, 
 which occur when multiple goroutines access shared resources simultaneously, 
 leading to inconsistent or incorrect data.*/

 
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


var wg = sync.WaitGroup{} 
var mu = sync.Mutex{}//creating a mutex
 

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var result = []string{}



func main() {
	t0 := time.Now()
	
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	
	wg.Wait() 
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("%v ",result)
}


func dbCall(i int) {
 	 
	delay := rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)

	 
	
	fmt.Printf("The result from the database is: %s\n", dbData[i])


	// Lock the mutex before accessing the shared resource
	mu.Lock()
	result = append(result,dbData[i])
	// Unlock the mutex after accessing the shared resource
	mu.Unlock()

	defer wg.Done() 
}





  

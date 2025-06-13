 /* this code will show the execuation without the concurrency, I mean how the code is going the
 work without concurrency*/ 
 package main

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
}

func dbCall(i int) {
	// Simulate DB call delay
	delay := rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
}



/*Since the database calls are made sequentially (one after the other), 
the total execution time will be the sum of all individual delays, 
which can be significant if the delays are long.*/
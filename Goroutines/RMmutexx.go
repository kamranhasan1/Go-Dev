/*with Read and Write mutex
The RLock is used for reading, and WLock is used for writing.
 This allows multiple readers to access the result slice simultaneously 
 while ensuring exclusive access for writers.
*/


package main

import (
	"fmt"
	// "log"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var mu = sync.Mutex{}  
var m = sync.RWMutex{} 

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
	fmt.Printf("\nthe res are %v ", result)
}

func dbCall(i int) {

	delay := rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()

}

func save(res string) {
	m.Lock()
	result = append(result, res)
	m.Unlock()

}

func log() {
	m.RLock()
	fmt.Printf("\nthe results are %v", result)
	m.RUnlock()
}

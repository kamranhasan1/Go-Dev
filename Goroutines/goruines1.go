package main

import (
	"fmt"
	"sync"
)
// var wg sync.WaitGroup

func work(id int,w *sync.WaitGroup){
	defer w.Done()
	fmt.Println("worker with ", id," is working")
}
func main(){
	var wg sync.WaitGroup
 
	for i := 0; i < 10 ; i++{
		wg.Add(1)
		go work(i, &wg) 
	}
	wg.Wait()
}
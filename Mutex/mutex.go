 //In simple words, we use a mutex to prevent race conditions while a program is running concurrently.
 package main

import (
	"fmt"
	"sync"
)

type post struct{
	views int
	like int 
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup){
	defer func (){
	defer wg.Done() 
	p.mu.Unlock()//2. I'm done, next person can go
	}()
	 

	p.mu.Lock()// 1. I'm using views now, wait your turn
	p.views += 1 
	p.like += 1 
 
}


func main(){
	var wg sync.WaitGroup

	myPost := post{views: 0,like:2}

	for i:=0; i<100; i++{
		wg.Add(1) 
		go myPost.inc(&wg)

	}
	 wg.Wait()
	fmt.Println("likes: ",myPost.like,"views",myPost.views)
}
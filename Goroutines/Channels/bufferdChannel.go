package main

import (
	"fmt"
	"time"
)

//bufferd channel
func main(){
	
	emailChan := make(chan string, 100)
	done := make(chan bool)
	go emailSender(emailChan,done)

	for i := 0 ; i<100; i++{
		emailChan <- fmt.Sprint("%d@gmail.com",i)
	}
	fmt.Println("done sending")
	 
	<-done
} 



func emailSender(emailChan chan string,done chan bool){
	
	defer func ()  {done <- true}()

	for email:= range emailChan{
		fmt.Println("sending email to ",email)
		time.Sleep(time.Second *1)
	}
}
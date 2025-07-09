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

	for i := 0 ; i<2; i++{
		emailChan <- fmt.Sprint("%d@gmail.com",i)
	}
	fmt.Println("done sending")
	//this is important to close 
	close(emailChan)
	 
	<-done
	
	//double channel 
	chan1 := make(chan string)
	chan2 := make(chan int)
	go data(chan1,chan2)
	time.Sleep(time.Second*1)
	



} 



func emailSender(emailChan chan string,done chan bool){
	
	defer func ()  {done <- true}()

	for email:= range emailChan{
		fmt.Println("sending email to ",email)
		time.Sleep(time.Second *1)
	}
}

//double channel 
func data(chan1 chan string, chan2 chan int ){
 
	go func ()  {
		chan1<-"ping pong "
		
	}()
	go func ()  {
		chan2<- 14
	}()

	for i:= 0; i<2; i++{
	select{
	case chan1Val:=<-chan1:
		fmt.Println("recived data from chan1",chan1Val)
	case chan2Val := <-chan2:
		fmt.Println("he was saying me ",chan2Val)
	}	
	}

}
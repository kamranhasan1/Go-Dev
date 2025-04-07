package main

import (
	"fmt"
	"math/rand"
	"time"
)

var maxChickenPrice float32 = 5
var maxPaneerPrice float32 = 3

func main() {
	t:= time.Now()
	
	var chickenChannel = make(chan string)
	var paneerChannel = make(chan string)
	var websites = []string{"Blinkit.com", "Zepto.com", "InstaMart.com"}
	for i := range websites {
		go checkChickenPrice(websites[i], chickenChannel)
		go checkPaneerPrice(websites[i],paneerChannel)
	}
	sendMessage(chickenChannel,paneerChannel)
	

	fmt.Printf("Total execution time: %v\n", time.Since(t))}

func checkChickenPrice(website string, chickenChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var chickenPrice = rand.Float32()*20 // Generates random number between 0-19
		if chickenPrice <= maxChickenPrice {
			chickenChannel <- website
			break
		}
	}
}


func checkPaneerPrice(website string,paneerChannel chan string){
	for{
		time.Sleep(1*time.Second)
		var paneerPrice  = rand.Float32()*20
		if paneerPrice< maxPaneerPrice{
			paneerChannel<-website
			break
		}
	}
}

func sendMessage(chickenChannel chan string,paneerChannel chan string) {
	select{
	case website := <-chickenChannel:
		fmt.Printf("Found a chicken deal on %s\n", website)
	case website:= <-paneerChannel:
		fmt.Printf("Found a Paneer deal on %s\n", website)

	}
}
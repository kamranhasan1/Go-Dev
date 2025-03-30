//more defined way
package main

import "fmt"

// Define a struct 'brainrot' with fields for reels, shorts, followers, and userInfo.
type brainrot struct {
	reels     uint8
	shorts    uint8
	followers uint16
	userInfo  user
}

 type user struct {
	name string
}


func main() {
	 
	var myFeed1 brainrot = brainrot{
		reels: 14, 
		shorts: 56,
		userInfo: user{name: "Aryan"}}

	    myFeed1.followers = 224

 	var myFeed brainrot = brainrot{12, 
		25, 
		30, 
		user{name: "satlxy_k"}} // Order: reels, shorts, followers, userInfo





 	fmt.Println(myFeed.reels, myFeed.shorts, myFeed.userInfo.name, myFeed.followers)  
	fmt.Println(myFeed1.reels, myFeed1.shorts, myFeed1.userInfo.name, myFeed1.followers) 
}


 
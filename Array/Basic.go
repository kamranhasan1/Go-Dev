package main
import (
	"fmt"	
)

func main(){
	//explicitly, size arr
	arr := [...]int32{2,3,4}
	fmt.Println(arr)

	//printing whole arr using loop // this is slice because it has no fixed size like vector in cpp
	var arrr []int32  = [] int32{23,34}
	for i:=0; i<len(arrr); i++{
		fmt.Println(arrr[i])
	 }

	 for index, value := range arrr {
		fmt.Println(index,value) // Print the value at the current index
	}
	

	 fmt.Println("capacity of array",cap(arrr))


	//printing address
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])

	
}




 


















 
 





 
package main 
import("fmt")
/*
can use "any" to implement generic

func printSlice[T any ](items []T){

	for _, item := range items {
		fmt.Println(item)


	}

}*/


/*



or "interface{}" will also work 

func printSlice[T interface{} ](items []T){

	for _, item := range items {
		fmt.Println(item)


	}

}
*/



// or "T int|string" can also be used 

func printSlice[T int|string ](items []T){

	for _, item := range items {
		fmt.Println(item)


	}

}

func main (){
	nums := []int{1,23,4,4}
	strs:=[]string{"what's","good"}
	printSlice(strs)
	printSlice(nums)
}
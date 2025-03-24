package main
import "fmt"
import "time"

func main() {
	 
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) 

	myMapp := map[string]uint8{"kamran": 23, "hasan": 34}
	fmt.Println(myMapp["kamran"])  



	// 3. Accessing a non-existent key returns the zero value (0 for uint8)
	fmt.Println(myMapp["ak"]) 



	// 4. Check if a key exists using `ok`
	var age, ok = myMapp["jhon"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid Name")  
	}




	// 5. Iterate over map using range loop
	var t0 = time.Now() // Record the start time
    for key, val := range myMapp {
        fmt.Printf("Name: %v, Age: %v\n", key, val) 
        fmt.Println("Elapsed time:", time.Since(t0)) // Print elapsed time since t0
    }
}

 
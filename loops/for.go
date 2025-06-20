package main
import ("fmt")

func main() { 
//while loop
 i := 1
 for i<=3 {
	fmt.Println(i)
	i= i+1

 //classic for loop

for i:= 0 ; i<3; i++{
	fmt.Println(i)

}
 

//using range keyword 
for i:= range 10{
	if i ==8{
		continue
	}
	fmt.Println(i)
}
 }}


//a variadic function is a function that can accept a variable number of arguments of the same type.

package main
import(
	"fmt"
)

func sum (nums ...int) int{//You define a variadic function using ... before the type of the last parameter.
	s := 0 
	for _,num := range nums{
		s = s+num
	}
	return s
}
func main(){
	res := sum(3,4,5,6,7,7,8)//here is the example
	numss := []int{3,4,5,6}
	res2 := sum(numss ...)
	fmt.Println(res2)//or we can do this also f
	fmt.Println(res)
}
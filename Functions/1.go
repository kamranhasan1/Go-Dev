package main
import(
	"fmt"
)

func add(a , b int ) int{
	return a*b
}

func getLang()(string, string, string){
	return "goLang","javaScript","C"
}
func main (){

	result:= add(3,4)
	fmt.Println(result)
	fmt.Println(getLang())
}
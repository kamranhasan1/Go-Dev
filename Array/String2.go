package main
import (
	"fmt"
	"strings"
)

func main() {
	// Your original concatenation
	StrSlice := []string{"k", "a", "m", "r", "a", "n"}
	emptyStr := ""
	for i := range StrSlice {
		emptyStr += StrSlice[i]
	}
	fmt.Printf("%v\n", emptyStr)   

	// Call the Sbuilder()
	Sbuilder()
}


// OR Using strings.Builder  
func Sbuilder() {
	StrSlices := []string{"k", "a", "m", "r", "a", "n"}
	var StrBuilder strings.Builder  //  Create empty strings.Builder object

	for i := range StrSlices {
		StrBuilder.WriteString(StrSlices[i]) // Append current string to builder
	}
	catStr := StrBuilder.String()  
	fmt.Print(catStr)  
}
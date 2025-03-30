package main
import "fmt"

type Animal struct {
    Dog
    Cat
}

type Dog struct {
    D_sound string
    D_Breed string
}

type Cat struct {
    C_sound string
    C_Breed string
}

func main() {
    var A Animal = Animal{
        Dog: Dog{
            D_sound: "Woof or Ruff",
            D_Breed: "Dachshunds",
        },
    }

    var B Animal = Animal{
        Cat: Cat{
            C_sound: "meow",
            C_Breed: "Scottish Fold",
        },
    }

    fmt.Println("Dog Sound- ",A.D_sound,", ","Dog Breed- ", A.D_Breed)  
    fmt.Println("Cat Sound- ",B.C_sound,", ","Cat Breed-",  B.C_Breed) 
}
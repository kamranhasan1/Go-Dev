package main
import ("fmt"
		"time")

// Define the Animal interface
type Animal interface {
    MakeSound() string
    GetBreed() string
    Describe() 
}

// Dog implements Animal
type Dog struct {
    sound string
    breed string
}

func (d Dog) MakeSound() string {
    return d.sound
}

func (d Dog) GetBreed() string {
    return d.breed
}

func (d Dog) Describe() {
    fmt.Printf("I'm a %s dog. I say %s!\n", d.breed, d.sound)
}

func (d Dog) WagTail() {
    fmt.Println("Dog is wagging its tail happily!")
}

// Cat implements Animal
type Cat struct {
    sound string
    breed string
}

func (c Cat) MakeSound() string {
    return c.sound
}

func (c Cat) GetBreed() string {
    return c.breed
}

func (c Cat) Describe() {
    fmt.Printf("I'm a %s cat. I say %s!\n", c.breed, c.sound)
}

func (c Cat) Purr() {
    fmt.Println("Cat is purring contentedly!")
}



func main() {

	t:= time.Now()
     
    animals := []Animal{
        Dog{sound: "Woof", breed: "Dachshund"},
        Cat{sound: "Meow", breed: "Scottish Fold"},
    }

    // Process all animals polymorphically
    for _, animal := range animals {
        animal.Describe()
        fmt.Println("Sound:", animal.MakeSound())
        fmt.Println("Breed:", animal.GetBreed())
        
        // Type assertion for specific behaviors
        switch a := animal.(type) {
        case Dog:
            a.WagTail()
        case Cat:
            a.Purr()
        }

		
         
        fmt.Println("-----")
		
    }

	fmt.Printf("\nProgram execution took %s\n", time.Since(t))

}
package main

import "fmt"

type Dog struct {
	sound string
	Breed string
}

type Cat struct {
	sound string
	Breed string
}

type Animal struct {
	Name string
	Dog  Dog
	Cat  Cat
}

// method to print their description
func (e Animal) PetName() {
	fmt.Printf("\nName: %s\n", e.Name)

	if e.Dog.Breed != "" {
		fmt.Printf("Type: Dog\nSound: %s\nBreed: %s\n", e.Dog.sound, e.Dog.Breed)
	}

	if e.Cat.Breed != "" {
		fmt.Printf("Type: Cat\nSound: %s\nBreed: %s\n", e.Cat.sound, e.Cat.Breed)
	}
}

func main() {
	var A Animal = Animal{
		Name: "Dog the Culo",
		Dog: Dog{
			sound: "Woof or Ruff",
			Breed: "Dachshunds",
		},
	}

	var B Animal = Animal{
		Name: "Cat the Billi",
		Cat: Cat{
			sound: "meow",
			Breed: "Scottish Fold",
		},
	}

	// calling the method
	A.PetName()
	B.PetName()
}

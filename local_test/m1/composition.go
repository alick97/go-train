package m1

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println("Generic animal sound")
}

type Animal2 struct {
	Name string
}

func (a *Animal2) Speak() {
	fmt.Println("Generic animal2 sound")
}

type Flyer struct {
	Wingspan int
}

func (f *Flyer) Fly() {
	fmt.Println("Flying...")
}

// Anonymous Embedding
type Bird struct {
	Animal
	Flyer
}

type BirdOverMethod struct {
	Animal
	Flyer
	Animal2
}

func CompositionExample() {
	myBird := Bird{
		Animal: Animal{Name: "Eagle"},
		Flyer:  Flyer{Wingspan: 200},
	}
	myBird.Speak()           // Output: Generic animal sound
	myBird.Fly()             // Output: Flying...
	fmt.Println(myBird.Name) // Output: Eagle
	myBird2 := BirdOverMethod{
		Animal:  Animal{Name: "Eagle"},
		Flyer:   Flyer{Wingspan: 200},
		Animal2: Animal2{Name: "Eagle2"},
	}
	myBird2.Animal.Speak()                                 // Output: Generic animal sound
	myBird2.Animal2.Speak()                                // Output: Generic animal sound
	myBird2.Fly()                                          // Output: Flying...
	fmt.Println(myBird2.Animal.Name, myBird2.Animal2.Name) // Output: Eagle
}

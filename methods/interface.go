package methods

import "fmt"

// The purpose of this program is to have a basic  understanding of does we mean by interface

// Declaring a type Name with field first and last
type Name struct {
	Name string
	Age  int
}

// Declaring a type Status with field inSameCity if true then we are in same city and Name
type Status struct {
	Name       string
	InSameCity bool
}

// Declaring a type Friend interface, Which means any type that is using the function speak is also of type Friend
type Friend interface {
	Speak()
}

// This is recving values for the type Name
func (n Name) Speak() {
	fmt.Println("Name isn called", n)
}

func (s *Status) Speak() {
	fmt.Println("Status is called")
}

func CallingInterfaceFriend(f Friend) {
	fmt.Println("I have passed an interface and value for that is: ", &f, f)
	fmt.Printf("The type of passed argument is: %T\n", f)
}

func PlayingWithInterface() {
	n1 := Name{
		Name: "Kushagra",
		Age:  26,
	}

	s1 := Status{
		Name:       "Pranav",
		InSameCity: true,
	}

	CallingInterfaceFriend(n1)
	CallingInterfaceFriend(&s1)
}

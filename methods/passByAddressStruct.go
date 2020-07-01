package methods 

import "fmt"

// The purpose of this program is to learn pass by address for struct

// Declaring a user defined Struct name User
type User struct{
	Name string 
	ID int
}

func UnderstandingPassByAddress(){
	u1 := User{
		Name : "Kushagra",
		ID : 007,
	}
	u2 := &User{}
	// I have multiple questions here
	// ? Question 1: & sign is used to access address then what is the difference between delcaration 1 and declaration2
	fmt.Println("\nAddress for both u1 and u2", &u1 , u2)	
	// * If we are printing the value the expectation is to get and address but instead of that I am getting the struct values (Default if not provided)

	// But if I am typing this I am getting the address
	fmt.Println("Memory address for u2: " , &u2)
	// And this whole blog for me is confusing from line 14 to 26

	
	ChangeUserName(u1)
	fmt.Println(u1)
	ChangeUserNameByPointer(u2) // ! For this it shows error ChangeUserName when we are passing address for struct and it should of type *User
	fmt.Println(&u2)
}

func ChangeUserName(u User){
	u.Name = "Kushgi"
	fmt.Println("Value of user is ch anged to: " , u.Name)
}

func ChangeUserNameByPointer(u *User){
	u.Name = "John"
}
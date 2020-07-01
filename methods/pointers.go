package methods

import "fmt"

// The purpose of this program is to understand what actual use of pointers is inm Go
func UnderstandingPointer() {
	// We will be declaring three variables out two variables will be of type pointer (*Type)
	// And we will change the value of variable two

	i := 10 // Short hand declaration for variable i because we are not bother about the int type means int, int8
	// var i int = 10 also valid declaration
	fmt.Println("\nValue of i: ", i)

	//i1 *int := &i We can't declare this it is not valid because the short hand declaration did not expect the type to be declared,
	// It will just check for identifer name to the left side

	i1 := &i
	changeValue(i1)
	fmt.Println("Value of i: ", i)
	//var i1 *int = &i
	// i1 := &i and var i1 *int = &i both are valid declaration because if we are not specifying the type it will take on this own
	// if *int is not defined then it will throw error reason because it is looking for a refrence to connect to that location

	// const i2 *int = i1 This is not valid  the reason being value of i1 can be changed and the we are assigning the address for i1
	// Which will allow you to change the value at position
	// Also if the variable is declared as constant you can't accerss the address

	i2 := i1
	// We can declare this as var i2 *int the explation is same for this also
	fmt.Println(i, &i, i1, &i1, *i1, i2, &i2, *i2)
	changeValue(i2)
	fmt.Println("Value of i: ", i)
}

// This won't be valid if we are passing i2 to this function because it is expecting a type *type (e.g *int)
// func changeValue(i int) {
// 	fmt.Println(i)
// }

func changeValue(a *int) {
	fmt.Println("Passed argument values: ", a, &a, *a)
	*a = *a + 10
}

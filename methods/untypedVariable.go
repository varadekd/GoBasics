package methods

import (
	"fmt"
)

// The purpose of this program is to check wether the untyped works for variable
// func UndeclaredTypeVar(val float32) {
// 	fmt.Printf("We got value %f of type %T", val, val)
// }

// This program will not work because the untyped and typed works
// In this function we are trying to check that we are passing int value
// And it is getting converted into float
func UndeclaredTypeConst(val float32) {
	fmt.Printf("\nWe got variable2 value %f of type %T", val, val)
}

// The purpose of this priogram is to check what is the different default type
// is used for the value declared
func UntypedDeclaration() {
	// Declaring short hand declaration for int, string , float and bool
	shortHandInt := 19
	shortHandFloat := 19.19
	shortHandBool := false
	shortHandString := "This is short hand declaration"

	// Printing the type and size of each variable
	fmt.Println("Short hand deation default type")
	fmt.Printf("\n%T\n%T\n%T\n%T\n", shortHandInt, shortHandFloat, shortHandString, shortHandBool)

	var varInt = 19
	var varFloat = 19.19
	var varBool = false
	var varString = "This is varaible declaration"

	// Printing the type and size of each variable
	fmt.Println("Varaible declaration default type")
	fmt.Printf("\n%T\n%T\n%T\n%T\n", varInt, varFloat, varBool, varString)

	const constInt = 19
	const constFloat = 19.19
	const constBool = false
	const constString = "This is constant declaration"

	// Printing the type and size of each variable
	fmt.Println("constant declaration default type")
	fmt.Printf("\n%T\n%T\n%T\n%T\n", constInt, constFloat, constBool, constString)
}

// We will use untyped declaration in cases where we are not bother about the which size we want to use.
// Like if we are not specifying var a = 10 then default type is int

package main

import (
	"fmt"

	"../src/methods"
)

var globalVariable int = 10

//var globalVariable string = 90
// The above declartion will not work since we are explicitly mentioned that indentifer should of type string and we are passing value int

// var globalVariable = 70 This is also a correct way fo declaring the variable here we are not specifying the type hence it will detect on our own

// globalVariable := 90 Why this declaration is not valid?

// a , b := 1, 0 This is a valid declaration and if we use it it won't show any error
// But
// a, b := 1 , a

func main() {
	// Printing message for package is called
	methods.PackageImportMessage()
	// methods.packageImportMessage(globalVariable) this will throw error saying too many aruguments are passed. Because function is
	// No arguments and vice versa is also possible that when function is expecting something and we are not passing anything

	// Why this throws an error
	// const constantValue = "I am string"
	// methods.ChangeStringValue(&constantValue)

	// fmt.Println("Prinitng the value of divisor: " , divisor)
	// Above syntax is not valid since the variable scope start after it is declared

	const divisor = 2
	// However this is valid
	fmt.Println("Prinitng the value of divisor: ", divisor)

	// This is a wrong way to access the package methods
	// methods.packageImportMessage()
	// Reason is we are trying to access the un exported function

	fmt.Println("\nValue of globalVariable before calling NoChangeInValue : ", globalVariable)
	methods.NoChangeInValue(&globalVariable)
	fmt.Println("Value of globalVariable before calling NoChangeInValue : ", globalVariable)
	// The reason in Go for every new variable that is declared a dedicated for variable so when we are assinging and then changing
	// The value of new variable so when we changed the new value it will relfect/change the value for that particular address

	fmt.Println("\nValue of globalVariable before calling ChangeValue: ", globalVariable)
	// I am passing the address where the globalVariable is sotered inside the memory.
	// If we are passing anything other than type int it will throw error reason being we are expecting the value to of type int
	methods.ChangeValue(&globalVariable)
	fmt.Println("Value of globalVariable after calling ChangeValue: ", globalVariable)
	// Passing the address will make sense only in the cases where I want to change the value stored on that address.

	// Example to access the value and then returing the computed values
	fmt.Println("\nWill return the sqaure of the value passed: ", methods.GetSqaure(globalVariable))

	// However changing value of the variable is also possible by returning a value from the function and then assinging it
	globalVariable = methods.ChangeValueThroughReturn()
	fmt.Println("\nValue of globalVariable after calling ChangeValueThroughReturn: ", globalVariable)
	// globalVariable := methods.ChangeValueThroughReturnWithValue(globalVariable) Explation is file
	// Why this cannot be used here

	// Functions returning multiple things
	remainder, err := methods.GetDivisionValue(globalVariable, divisor)
	fmt.Println("The value after divison is: ", remainder, err)

	// globalVariable = methods.ChangeValueThroughReturnString()
	// This will throw an error becase though the function is returning the valid value but the variable has a value of type int

	// The below part is valid since we are declaring a new variable and in case it won't show any error
	value := methods.ChangeValueThroughReturnString()
	fmt.Println("\nChangeValueThroughReturnString returned the value: ", value)

	// a := 7
	// fmt.Println("Value of a before calling the function: ", a)
	// methods.TestFunction(&a)
	// fmt.Println("Value of a before calling the function: ", a)
}

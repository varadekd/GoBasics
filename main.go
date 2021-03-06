package main

import (
	"../src/methods"
)

// If you are declaring the const for variable it should be done on package level
var globalVariable int = 10

//var globalVariable string = 90
// The above declartion will not work since we are explicitly mentioned that indentifer should of type string and we are passing value int

// var globalVariable = 70 This is also a correct way fo declaring the variable here we are not specifying the type hence it will detect on our own

// globalVariable := 90 Why this declaration is not valid?

// a , b := 1, 0 This is a valid declaration and if we use it it won't show any error
// But
// a, b := 1 , a

// ? Variable declaration valid types
// var a int is valid in this the default value of int is assigned to variable a
// var a int = 9 this is valid in this we are specifying that the variable is of type int and assigning the value 9 to it
// For const you need to provide value because language declaration looks for the value to be assigned to const variable // ? Why it is not taking default value
// But for short hand declaration that is not the case because it's defination says that you have to declare and assign value at the same time
//  So a := 1 is valid but a := int is not niether a := var b int, Reason if it is expecting the value in both cases it will throw error (Both error will be different)

// TODO: Ask question with Om regarding the type float and int from Question.go file - answer found

func main() {
	// // Printing message for package is called
	// methods.PackageImportMessage()
	// // methods.packageImportMessage(globalVariable) this will throw error saying too many aruguments are passed. Because function is
	// // No arguments and vice versa is also possible that when function is expecting something and we are not passing anything

	// // Why this throws an error
	// // const constantValue = "I am string"
	// // methods.ChangeStringValue(&constantValue)

	// // fmt.Println("Prinitng the value of divisor: " , divisor)
	// // Above syntax is not valid since the variable scope start after it is declared

	// // Type issue (Question to be asked) Answer found
	// // floatVal := 1
	// // methods.PassingFloatValue(floatVal)
	// // This wont work because the type of floatVal is int
	// // But if I use
	// // const floatVal = 1
	// // methods.PassingFloatValue(floatVal)
	// // Or methods.PassingFloatValue(1)
	// // This will work because for const value the context is changed since it is untyped

	// // Creating methods with same name and different signature
	// // methods.PrintPassedValue(10)
	// // methods.PrintPassedValue(10.10)

	// const divisor = 2
	// // However this is valid
	// fmt.Println("Prinitng the value of divisor: ", divisor)

	// // This is a wrong way to access the package methods
	// // methods.packageImportMessage()
	// // Reason is we are trying to access the un exported function

	// fmt.Println("\nValue of globalVariable before calling NoChangeInValue : ", globalVariable)
	// methods.NoChangeInValue(&globalVariable)
	// fmt.Println("Value of globalVariable before calling NoChangeInValue : ", globalVariable)
	// // The reason in Go for every new variable that is declared a dedicated for variable so when we are assinging and then changing
	// // The value of new variable so when we changed the new value it will relfect/change the value for that particular address

	// fmt.Println("\nValue of globalVariable before calling ChangeValue: ", globalVariable)
	// // I am passing the address where the globalVariable is sotered inside the memory.
	// // If we are passing anything other than type int it will throw error reason being we are expecting the value to of type int
	// methods.ChangeValue(&globalVariable)
	// fmt.Println("Value of globalVariable after calling ChangeValue: ", globalVariable)
	// // Passing the address will make sense only in the cases where I want to change the value stored on that address.

	// // globalVariable = methods.ChangeValueThroughReturnString()
	// // This will throw an error becase though the function is returning the valid value but the variable has a value of type int

	// // The below part is valid since we are declaring a new variable and in case it won't show any error
	// value := methods.ChangeValueThroughReturnString()
	// fmt.Println("\nChangeValueThroughReturnString returned the value: ", value)

	// // Example to access the value and then returing the computed values
	// fmt.Println("\nWill return the sqaure of the value passed: ", methods.GetSqaure(globalVariable))

	// // However changing value of the variable is also possible by returning a value from the function and then assinging it
	// globalVariable = methods.ChangeValueThroughReturn()
	// fmt.Println("\nValue of globalVariable after calling ChangeValueThroughReturn: ", globalVariable)
	// // globalVariable := methods.ChangeValueThroughReturnWithValue(globalVariable) Explation is file
	// // Why this cannot be used here

	// // Functions returning multiple things
	// remainder, err := methods.GetDivisionValue(globalVariable, divisor)
	// fmt.Println("The value after divison is: ", remainder, err)

	// // Performs all the basic arthimetic operations
	// methods.BasicArthimeticOperations()

	// // Functions for checking typed and untyped
	// methods.UntypedConst(2.1)
	// // methods.TypedConst(2.1)

	// // Below functions are checking the declaration of boolean
	// methods.DeclareBoolTrue()
	// methods.DeclareBoolFalse()
	// methods.DeclareBoolFromCondition(90, 89)

	// // Conversion when untyped variable is passed
	// // var variable1 = 90
	// // methods.UndeclaredTypeVar(variable1)
	// const variable2 = 90
	// fmt.Printf("\nValue %d of variable2 before calling UndeclaredTypeConst and its type is %T", variable2, variable2)
	// methods.UndeclaredTypeConst(variable2)

	// methods.UntypedDeclaration() // Will print all the default types for the variable declaration

	// // The puspose of this is to learn about arthimetic operations on pointer
	// // val1 := 10
	// // methods.OperationsOnPointer(&val1)
	// // Explanation is inside arthimeticOperation.go

	// // const variable3 float32 = 90.0
	// // fmt.Printf("\nValue %d of variable3 before calling UndeclaredTypeConst and its type is %T", variable3, variable3)
	// // methods.DeclaredTypeConst(variable3)

	// // Learning of pointer the below methods call the UnderstandingPointer function which further calls changePointerValue function
	// // Check file pointers.go
	// methods.UnderstandingPointer()
	// methods.UnderstandingPointerPointer()

	// // Using user defined type
	// //  Valid type of declaration
	// //  var person structs.Person we have delcared a variable person of type Person, Currently all the values are set to default
	// //  var person structs.Person = structs.Person{}  This is same as above because we have assigned the structs.Person with default values
	// //  var person = structs.Person{}  This will detect the variable is of type Person on its own.

	// person := structs.Person{
	// 	Firstname: "",
	// 	Lastname:  "Varade",
	// }

	// fmt.Printf("\nInital name: %+v\n", person)
	// methods.AccessPerson(person)
	// // If we pass like this the address for the variable decalred and varible passed will be different
	// // Means if we print the address of person here and the passed argument address inside the function they both will differ
	// fmt.Printf("Value of person after calling accessPerson %+v\n", person)
	// methods.ChangeFirstName(&person)
	// fmt.Printf("Value of person after calling ChangeFirstName %+v\n", person)

	// // Exploring more about the pass by address for the user defined type
	// methods.UnderstandingPassByAddress()

	// Playing with interfaces
	methods.PlayingWithInterface()
}

package methods

import (
	"errors"
	"fmt"
)

//  The purpose of this program is to get the arthimetic operation and return its type
func GetDivisionValue(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("The value of divisor cannot be zero")
	}

	return (dividend / divisor), nil
	// It is necessary for you to return the value since the function is expecting it to return something
}

// Written in doubts
func GetMultiplication(multiplicant1 float32, multiplicant2 float32) {
	fmt.Println(multiplicant1 * multiplicant2)
}

func BasicArthimeticOperations() {
	fmt.Println("\n", 1+1.9)
	fmt.Println(1 * 1.9)
	fmt.Println(3.9 - 3)
	fmt.Println(3.9 / 3)
	// fmt.Println("22" - "2")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	fmt.Println(!true)
	fmt.Println(!false)
}

// const a = 4.0
// 	fmt.Printf("%T%T%T ", a, multiplicant1, multiplicant2)
// 	fmt.Println("Multiplication of a float value and int is: ", (a*multiplicant1)+(a*multiplicant2))
// 	fmt.Println(multiplicant1 / a)

// Special notes when working with const in GO
//  const in go are typed or untyped and for the untyped the value gets change if the context is changed.
// Example
// const a = 1 is of type int
// But if it is used like this 2.1 * 1 it will act as float64 (The default type for untyped const if float64 for floating point)
// For the below program is valid
func UntypedConst(val float32) {
	const a = 2
	fmt.Println("Printing multiplication of untyped const and val passed as argument: ", val*a)
}

// But the below program will not work
// func TypedConst(val float32) {
// 	const a int = 2
// 	fmt.Println("Printing multiplication of typed const and val passed as argument: ", val*a)
// }

// The below functions is to perform arthimetic operations on pointer type
// func OperationsOnPointer(a *int) {
// 	fmt.Println("Value of a inside OperationsOnPointer: ", a)
// 	a = a + a // Any arthimetic operation is not valid for the pointer type
// }

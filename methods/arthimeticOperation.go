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
// func GetMultiplication(multiplicant1 float32, multiplicant2 int) {
// 	fmt.Println(multiplicant1 * multiplicant2)
// }

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

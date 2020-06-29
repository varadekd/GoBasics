package methods

import (
	"errors"
	"fmt"
)

func PackageImportMessage() {
	fmt.Println("Package components are properly imported")
}

// func packageImportMessage() {
// 	fmt.Println("Package components are properly imported")
// }
// The above function declaration is also valid but we can't use this outside the package since it is not exported
// How ever if you want to use it we can use in the way commented below
// func ChangeValue(a *int) {
// 	packageImportMessage()
// 	*a = 90
// }

func NoChangeInValue(a *int) {
	val := *a
	val = 80
	fmt.Println("Value changed to: ", val)
}

func ChangeValue(a *int) {
	// *a this helps us to us to access the value that is stored at the passed address
	*a = 90
}

func ChangeValueThroughReturn() int {
	return 231
}

func GetSqaure(val int) int {
	return val * val

	// Can be done in this fashion also but I will pefer the above way because In this context I have to return sqaure so not storing this value
	// newVal = val * val
	// return newVal
	// OR like this we arer squaring the passed val and storing the same inside the passed value and then returning it
	// val = val * val
	// return val
}

// This is also valid but I will not pefer this because my intention is to return value without performing any operation on the inital value
// So there is no point in passing the value where the motive is to change the value
// func ChangeValueThroughReturnWithValue(a int) int {
// 	a = 200
// 	return a
// }

// The below one is also value we are returning a value of type string
func ChangeValueThroughReturnString() string {
	return "This is a string"
}

func ChangeStringValue(s *string) {
	*s = "The string is changed"
}

// However this is not valid
// func ChangeValueThroughReturnString() string {
// 	return 230
// }
// Reason for this is we are saying that the function will return a value of type string and in actual we are returning a value of type int

func GetDivisionValue(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("The value of divisor cannot be zero")
	}

	return (dividend / divisor), nil
	// It is necessary for you to return the value since the function is expecting it to return something
}

// This is just test function
// func TestFunction(val *int) {
// 	*val = 9
// 	// 	fmt.Println("Printing values from test function: ", val)
// 	// 	var a string
// 	// 	return a
// 	// 	// return variable or return "somrthing is valid"
// 	// 	// But returning return var a string, return string is not valid
// }

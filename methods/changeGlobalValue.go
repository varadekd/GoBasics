package methods

import "fmt"

// The purpose of this program is to change the global value that is passed by address

func ChangeValue(a *int) {
	fmt.Println("Function Change value is triggered")
	// *a this helps us to us to access the value that is stored at the passed address
	*a = 90
}

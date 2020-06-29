package methods

import "fmt"

// The purpose of this program is to check if the argument is passed as addresss and is assigned to new variable the value should not change
func NoChangeInValue(a *int) {
	val := *a
	val = 80
	fmt.Println("Value changed to: ", val)
}

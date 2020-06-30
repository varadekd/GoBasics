package methods

import "fmt"

// The purpose of this program is to check wether the untyped works for variable
// func UndeclaredTypeVar(val float32) {
// 	fmt.Printf("We got value %f of type %T", val, val)
// }

// This program will not work because the untyped and typed works

func UndeclaredTypeConst(val float32) {
	fmt.Printf("\nWe got variable2 value %f of type %T", val, val)
}

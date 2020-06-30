package methods

import "fmt"

// The purpose of this program is to test the declaration of boolean

func DeclareBoolTrue() {
	a := true
	fmt.Println("\nPrinting not for variable a: ", !a)
}

func DeclareBoolFalse() {
	a := false
	fmt.Println("Value for a: ", a)
}

func DeclareBoolFromCondition(val1, val2 int) {
	a := val1 == val2
	fmt.Println("Value for a inside declaration of variable through condition: ", a)
}

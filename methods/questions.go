package methods

// ? @om Why this is valid in the below context _ question is answered by Kushagra
import "fmt"

func PassingFloatValue(a float32) {
	fmt.Printf("The value passed to PassingFloatValue is %f and it's type is %T", a, a)
}

package methods

import "errors"

//  The purpose of this program is to get the arthimetic operation and return its type
func GetDivisionValue(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("The value of divisor cannot be zero")
	}

	return (dividend / divisor), nil
	// It is necessary for you to return the value since the function is expecting it to return something
}

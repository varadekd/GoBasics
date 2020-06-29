package methods

// The pupose of this program is to access the global value and then return its sqaure
func GetSqaure(val int) int {
	return val * val

	// Can be done in this fashion also but I will pefer the above way because In this context I have to return sqaure so not storing this value
	// newVal = val * val
	// return newVal
	// OR like this we arer squaring the passed val and storing the same inside the passed value and then returning it
	// val = val * val
	// return val
}

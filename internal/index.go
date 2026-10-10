package internal

// Get_Index returns the element at index i from slice.
// If index is out of bounds, it returns the last element of the slice.
// If index < 0 , Get_Index returns ""
func Get_Index(slice []string, index int) string {

	var value string

	size := len(slice) - 1

	if index < 0 {

		return value

	}

	if size > index {

		value = slice[index]

	} else {

		value = slice[size]

	}

	return value

}

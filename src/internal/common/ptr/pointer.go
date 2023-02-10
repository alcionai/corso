package ptr

// ptr package is a common package used for pointer
// access and deserialization.

// func Value[T any](source T, ptr *T) (T, bool) {
// 	if ptr == nil {
// 		return source, false
// 	}

// 	return *ptr, true
// }

// Val helper method for unwrapping strings
// Microsoft Graph saves many variables as string pointers.
// Function will safely check if the point is nil prior to
// dereferencing the pointer. If the pointer is nil,
// an empty string is returned.
func Val(ptr *string) string {
	if ptr == nil {
		return ""
	}

	return *ptr
}

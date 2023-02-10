package ptr

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

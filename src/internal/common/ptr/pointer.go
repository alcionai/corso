package ptr

// ptr package is a common package used for pointer
// access and deserialization.

// Value Generic option for dereferencing pointers.
// Microsoft Graph saves many variables as string pointers.
// Function will safely check if the point is nil prior to
// dereferencing the pointer. If the pointer is nil,
// an empty version of the object is returned.
// Operation does not work on Nested objects.
// For example:
// *evt.GetEnd().GetDateTime() will still cause a panic
// if evt is nil or GetEnd() is nil
func Value[T any](ptr *T) T {
	if ptr == nil {
		return *new(T)
	}

	return *ptr
}

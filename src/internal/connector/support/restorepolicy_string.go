// Code generated by "stringer -type=RestorePolicy"; DO NOT EDIT.

package support

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Copy-1]
	_ = x[Drop-2]
	_ = x[Replace-3]
}

const _RestorePolicy_name = "UnknownCopyDropReplace"

var _RestorePolicy_index = [...]uint8{0, 7, 11, 15, 22}

func (i RestorePolicy) String() string {
	if i < 0 || i >= RestorePolicy(len(_RestorePolicy_index)-1) {
		return "RestorePolicy(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RestorePolicy_name[_RestorePolicy_index[i]:_RestorePolicy_index[i+1]]
}

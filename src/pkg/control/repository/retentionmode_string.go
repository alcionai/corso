// Code generated by "stringer -type=RetentionMode -linecomment"; DO NOT EDIT.

package repository

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownRetention-0]
	_ = x[NoRetention-1]
	_ = x[GovernanceRetention-2]
	_ = x[ComplianceRetention-3]
}

const _RetentionMode_name = "UnknownRetentionnonegovernancecompliance"

var _RetentionMode_index = [...]uint8{0, 16, 20, 30, 40}

func (i RetentionMode) String() string {
	if i < 0 || i >= RetentionMode(len(_RetentionMode_index)-1) {
		return "RetentionMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RetentionMode_name[_RetentionMode_index[i]:_RetentionMode_index[i+1]]
}

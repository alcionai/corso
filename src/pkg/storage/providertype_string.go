// Code generated by "stringer -type=ProviderType -linecomment"; DO NOT EDIT.

package storage

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ProviderUnknown-0]
	_ = x[ProviderS3-1]
	_ = x[ProviderFilesystem-2]
}

const _ProviderType_name = "Unknown ProviderS3Filesystem"

var _ProviderType_index = [...]uint8{0, 16, 18, 28}

func (i ProviderType) String() string {
	if i < 0 || i >= ProviderType(len(_ProviderType_index)-1) {
		return "ProviderType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ProviderType_name[_ProviderType_index[i]:_ProviderType_index[i+1]]
}

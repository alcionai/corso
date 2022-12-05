// Code generated by "stringer -type=CategoryType -linecomment"; DO NOT EDIT.

package path

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownCategory-0]
	_ = x[EmailCategory-1]
	_ = x[ContactsCategory-2]
	_ = x[EventsCategory-3]
	_ = x[FilesCategory-4]
	_ = x[ListsCategory-5]
	_ = x[LibrariesCategory-6]
}

const _CategoryType_name = "UnknownCategoryemailcontactseventsfileslistslibraries"

var _CategoryType_index = [...]uint8{0, 15, 20, 28, 34, 39, 44, 53}

func (i CategoryType) String() string {
	if i < 0 || i >= CategoryType(len(_CategoryType_index)-1) {
		return "CategoryType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CategoryType_name[_CategoryType_index[i]:_CategoryType_index[i+1]]
}

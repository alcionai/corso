// Code generated by "stringer -type=service -linecomment"; DO NOT EDIT.

package selectors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ServiceUnknown-0]
	_ = x[ServiceExchange-1]
	_ = x[ServiceOneDrive-2]
}

const _service_name = "Unknown ServiceExchangeOneDrive"

var _service_index = [...]uint8{0, 15, 23, 31}

func (i service) String() string {
	if i < 0 || i >= service(len(_service_index)-1) {
		return "service(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _service_name[_service_index[i]:_service_index[i+1]]
}

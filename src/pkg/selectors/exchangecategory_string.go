// Code generated by "stringer -type=exchangeCategory"; DO NOT EDIT.

package selectors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ExchangeCategoryUnknown-0]
	_ = x[ExchangeContact-1]
	_ = x[ExchangeContactFolder-2]
	_ = x[ExchangeEvent-3]
	_ = x[ExchangeMail-4]
	_ = x[ExchangeMailFolder-5]
	_ = x[ExchangeUser-6]
}

const _exchangeCategory_name = "ExchangeCategoryUnknownExchangeContactExchangeContactFolderExchangeEventExchangeMailExchangeMailFolderExchangeUser"

var _exchangeCategory_index = [...]uint8{0, 23, 38, 59, 72, 84, 102, 114}

func (i exchangeCategory) String() string {
	if i < 0 || i >= exchangeCategory(len(_exchangeCategory_index)-1) {
		return "exchangeCategory(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _exchangeCategory_name[_exchangeCategory_index[i]:_exchangeCategory_index[i+1]]
}

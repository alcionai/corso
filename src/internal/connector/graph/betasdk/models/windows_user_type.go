package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsUserType int

const (
    ADMINISTRATOR_WINDOWSUSERTYPE WindowsUserType = iota
    STANDARD_WINDOWSUSERTYPE
)

func (i WindowsUserType) String() string {
    return []string{"administrator", "standard"}[i]
}
func ParseWindowsUserType(v string) (interface{}, error) {
    result := ADMINISTRATOR_WINDOWSUSERTYPE
    switch v {
        case "administrator":
            result = ADMINISTRATOR_WINDOWSUSERTYPE
        case "standard":
            result = STANDARD_WINDOWSUSERTYPE
        default:
            return 0, errors.New("Unknown WindowsUserType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsUserType(values []WindowsUserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

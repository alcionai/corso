package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AndroidForWorkBindStatus int

const (
    NOTBOUND_ANDROIDFORWORKBINDSTATUS AndroidForWorkBindStatus = iota
    BOUND_ANDROIDFORWORKBINDSTATUS
    BOUNDANDVALIDATED_ANDROIDFORWORKBINDSTATUS
    UNBINDING_ANDROIDFORWORKBINDSTATUS
)

func (i AndroidForWorkBindStatus) String() string {
    return []string{"notBound", "bound", "boundAndValidated", "unbinding"}[i]
}
func ParseAndroidForWorkBindStatus(v string) (interface{}, error) {
    result := NOTBOUND_ANDROIDFORWORKBINDSTATUS
    switch v {
        case "notBound":
            result = NOTBOUND_ANDROIDFORWORKBINDSTATUS
        case "bound":
            result = BOUND_ANDROIDFORWORKBINDSTATUS
        case "boundAndValidated":
            result = BOUNDANDVALIDATED_ANDROIDFORWORKBINDSTATUS
        case "unbinding":
            result = UNBINDING_ANDROIDFORWORKBINDSTATUS
        default:
            return 0, errors.New("Unknown AndroidForWorkBindStatus value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkBindStatus(values []AndroidForWorkBindStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

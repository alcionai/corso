package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AndroidManagedStoreAccountBindStatus int

const (
    NOTBOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS AndroidManagedStoreAccountBindStatus = iota
    BOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
    BOUNDANDVALIDATED_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
    UNBINDING_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
)

func (i AndroidManagedStoreAccountBindStatus) String() string {
    return []string{"notBound", "bound", "boundAndValidated", "unbinding"}[i]
}
func ParseAndroidManagedStoreAccountBindStatus(v string) (interface{}, error) {
    result := NOTBOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
    switch v {
        case "notBound":
            result = NOTBOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
        case "bound":
            result = BOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
        case "boundAndValidated":
            result = BOUNDANDVALIDATED_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
        case "unbinding":
            result = UNBINDING_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
        default:
            return 0, errors.New("Unknown AndroidManagedStoreAccountBindStatus value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedStoreAccountBindStatus(values []AndroidManagedStoreAccountBindStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

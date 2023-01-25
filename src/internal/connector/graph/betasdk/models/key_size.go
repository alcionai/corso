package models
import (
    "errors"
)
// Provides operations to call the add method.
type KeySize int

const (
    // 1024 Bits.
    SIZE1024_KEYSIZE KeySize = iota
    // 2048 Bits.
    SIZE2048_KEYSIZE
    // 4096 Bits.
    SIZE4096_KEYSIZE
)

func (i KeySize) String() string {
    return []string{"size1024", "size2048", "size4096"}[i]
}
func ParseKeySize(v string) (interface{}, error) {
    result := SIZE1024_KEYSIZE
    switch v {
        case "size1024":
            result = SIZE1024_KEYSIZE
        case "size2048":
            result = SIZE2048_KEYSIZE
        case "size4096":
            result = SIZE4096_KEYSIZE
        default:
            return 0, errors.New("Unknown KeySize value: " + v)
    }
    return &result, nil
}
func SerializeKeySize(values []KeySize) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

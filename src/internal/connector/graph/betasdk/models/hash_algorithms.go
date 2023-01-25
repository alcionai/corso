package models
import (
    "errors"
)
// Provides operations to call the add method.
type HashAlgorithms int

const (
    // SHA-1 Hash Algorithm.
    SHA1_HASHALGORITHMS HashAlgorithms = iota
    // SHA-2 Hash Algorithm.
    SHA2_HASHALGORITHMS
)

func (i HashAlgorithms) String() string {
    return []string{"sha1", "sha2"}[i]
}
func ParseHashAlgorithms(v string) (interface{}, error) {
    result := SHA1_HASHALGORITHMS
    switch v {
        case "sha1":
            result = SHA1_HASHALGORITHMS
        case "sha2":
            result = SHA2_HASHALGORITHMS
        default:
            return 0, errors.New("Unknown HashAlgorithms value: " + v)
    }
    return &result, nil
}
func SerializeHashAlgorithms(values []HashAlgorithms) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

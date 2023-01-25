package models
import (
    "errors"
)
// Provides operations to call the add method.
type CertificateStore int

const (
    USER_CERTIFICATESTORE CertificateStore = iota
    MACHINE_CERTIFICATESTORE
)

func (i CertificateStore) String() string {
    return []string{"user", "machine"}[i]
}
func ParseCertificateStore(v string) (interface{}, error) {
    result := USER_CERTIFICATESTORE
    switch v {
        case "user":
            result = USER_CERTIFICATESTORE
        case "machine":
            result = MACHINE_CERTIFICATESTORE
        default:
            return 0, errors.New("Unknown CertificateStore value: " + v)
    }
    return &result, nil
}
func SerializeCertificateStore(values []CertificateStore) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

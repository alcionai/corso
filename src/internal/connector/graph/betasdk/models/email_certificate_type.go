package models
import (
    "errors"
)
// Provides operations to call the add method.
type EmailCertificateType int

const (
    // Do not use a certificate as a source.
    NONE_EMAILCERTIFICATETYPE EmailCertificateType = iota
    // Use an certificate for certificate source.
    CERTIFICATE_EMAILCERTIFICATETYPE
    // Use a derived credential for certificate source.
    DERIVEDCREDENTIAL_EMAILCERTIFICATETYPE
)

func (i EmailCertificateType) String() string {
    return []string{"none", "certificate", "derivedCredential"}[i]
}
func ParseEmailCertificateType(v string) (interface{}, error) {
    result := NONE_EMAILCERTIFICATETYPE
    switch v {
        case "none":
            result = NONE_EMAILCERTIFICATETYPE
        case "certificate":
            result = CERTIFICATE_EMAILCERTIFICATETYPE
        case "derivedCredential":
            result = DERIVEDCREDENTIAL_EMAILCERTIFICATETYPE
        default:
            return 0, errors.New("Unknown EmailCertificateType value: " + v)
    }
    return &result, nil
}
func SerializeEmailCertificateType(values []EmailCertificateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

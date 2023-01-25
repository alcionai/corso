package models
import (
    "errors"
)
// Provides operations to call the add method.
type CertificateValidityPeriodScale int

const (
    // Days.
    DAYS_CERTIFICATEVALIDITYPERIODSCALE CertificateValidityPeriodScale = iota
    // Months.
    MONTHS_CERTIFICATEVALIDITYPERIODSCALE
    // Years.
    YEARS_CERTIFICATEVALIDITYPERIODSCALE
)

func (i CertificateValidityPeriodScale) String() string {
    return []string{"days", "months", "years"}[i]
}
func ParseCertificateValidityPeriodScale(v string) (interface{}, error) {
    result := DAYS_CERTIFICATEVALIDITYPERIODSCALE
    switch v {
        case "days":
            result = DAYS_CERTIFICATEVALIDITYPERIODSCALE
        case "months":
            result = MONTHS_CERTIFICATEVALIDITYPERIODSCALE
        case "years":
            result = YEARS_CERTIFICATEVALIDITYPERIODSCALE
        default:
            return 0, errors.New("Unknown CertificateValidityPeriodScale value: " + v)
    }
    return &result, nil
}
func SerializeCertificateValidityPeriodScale(values []CertificateValidityPeriodScale) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CertificateRevocationStatus int

const (
    // Not revoked.
    NONE_CERTIFICATEREVOCATIONSTATUS CertificateRevocationStatus = iota
    // Revocation pending.
    PENDING_CERTIFICATEREVOCATIONSTATUS
    // Revocation command issued.
    ISSUED_CERTIFICATEREVOCATIONSTATUS
    // Revocation failed.
    FAILED_CERTIFICATEREVOCATIONSTATUS
    // Revoked.
    REVOKED_CERTIFICATEREVOCATIONSTATUS
)

func (i CertificateRevocationStatus) String() string {
    return []string{"none", "pending", "issued", "failed", "revoked"}[i]
}
func ParseCertificateRevocationStatus(v string) (interface{}, error) {
    result := NONE_CERTIFICATEREVOCATIONSTATUS
    switch v {
        case "none":
            result = NONE_CERTIFICATEREVOCATIONSTATUS
        case "pending":
            result = PENDING_CERTIFICATEREVOCATIONSTATUS
        case "issued":
            result = ISSUED_CERTIFICATEREVOCATIONSTATUS
        case "failed":
            result = FAILED_CERTIFICATEREVOCATIONSTATUS
        case "revoked":
            result = REVOKED_CERTIFICATEREVOCATIONSTATUS
        default:
            return 0, errors.New("Unknown CertificateRevocationStatus value: " + v)
    }
    return &result, nil
}
func SerializeCertificateRevocationStatus(values []CertificateRevocationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

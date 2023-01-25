package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CertificateDestinationStore int

const (
    // Computer Certificate Store - Root.
    COMPUTERCERTSTOREROOT_CERTIFICATEDESTINATIONSTORE CertificateDestinationStore = iota
    // Computer Certificate Store - Intermediate.
    COMPUTERCERTSTOREINTERMEDIATE_CERTIFICATEDESTINATIONSTORE
    // User Certificate Store - Intermediate.
    USERCERTSTOREINTERMEDIATE_CERTIFICATEDESTINATIONSTORE
)

func (i CertificateDestinationStore) String() string {
    return []string{"computerCertStoreRoot", "computerCertStoreIntermediate", "userCertStoreIntermediate"}[i]
}
func ParseCertificateDestinationStore(v string) (interface{}, error) {
    result := COMPUTERCERTSTOREROOT_CERTIFICATEDESTINATIONSTORE
    switch v {
        case "computerCertStoreRoot":
            result = COMPUTERCERTSTOREROOT_CERTIFICATEDESTINATIONSTORE
        case "computerCertStoreIntermediate":
            result = COMPUTERCERTSTOREINTERMEDIATE_CERTIFICATEDESTINATIONSTORE
        case "userCertStoreIntermediate":
            result = USERCERTSTOREINTERMEDIATE_CERTIFICATEDESTINATIONSTORE
        default:
            return 0, errors.New("Unknown CertificateDestinationStore value: " + v)
    }
    return &result, nil
}
func SerializeCertificateDestinationStore(values []CertificateDestinationStore) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

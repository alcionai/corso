package models
import (
    "errors"
)
// Provides operations to call the add method.
type VpnIntegrityAlgorithmType int

const (
    // SHA2-256
    SHA2_256_VPNINTEGRITYALGORITHMTYPE VpnIntegrityAlgorithmType = iota
    // SHA1-96
    SHA1_96_VPNINTEGRITYALGORITHMTYPE
    // SHA1-160
    SHA1_160_VPNINTEGRITYALGORITHMTYPE
    // SHA2-384
    SHA2_384_VPNINTEGRITYALGORITHMTYPE
    // SHA2-512
    SHA2_512_VPNINTEGRITYALGORITHMTYPE
    // MD5
    MD5_VPNINTEGRITYALGORITHMTYPE
)

func (i VpnIntegrityAlgorithmType) String() string {
    return []string{"sha2_256", "sha1_96", "sha1_160", "sha2_384", "sha2_512", "md5"}[i]
}
func ParseVpnIntegrityAlgorithmType(v string) (interface{}, error) {
    result := SHA2_256_VPNINTEGRITYALGORITHMTYPE
    switch v {
        case "sha2_256":
            result = SHA2_256_VPNINTEGRITYALGORITHMTYPE
        case "sha1_96":
            result = SHA1_96_VPNINTEGRITYALGORITHMTYPE
        case "sha1_160":
            result = SHA1_160_VPNINTEGRITYALGORITHMTYPE
        case "sha2_384":
            result = SHA2_384_VPNINTEGRITYALGORITHMTYPE
        case "sha2_512":
            result = SHA2_512_VPNINTEGRITYALGORITHMTYPE
        case "md5":
            result = MD5_VPNINTEGRITYALGORITHMTYPE
        default:
            return 0, errors.New("Unknown VpnIntegrityAlgorithmType value: " + v)
    }
    return &result, nil
}
func SerializeVpnIntegrityAlgorithmType(values []VpnIntegrityAlgorithmType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

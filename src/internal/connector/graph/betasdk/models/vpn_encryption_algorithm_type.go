package models
import (
    "errors"
)
// Provides operations to call the add method.
type VpnEncryptionAlgorithmType int

const (
    // AES-256
    AES256_VPNENCRYPTIONALGORITHMTYPE VpnEncryptionAlgorithmType = iota
    // DES
    DES_VPNENCRYPTIONALGORITHMTYPE
    // 3DES
    TRIPLEDES_VPNENCRYPTIONALGORITHMTYPE
    // AES-128
    AES128_VPNENCRYPTIONALGORITHMTYPE
    // AES-128-GCM
    AES128GCM_VPNENCRYPTIONALGORITHMTYPE
    // AES-256-GCM
    AES256GCM_VPNENCRYPTIONALGORITHMTYPE
    // AES-192
    AES192_VPNENCRYPTIONALGORITHMTYPE
    // AES-192-GCM
    AES192GCM_VPNENCRYPTIONALGORITHMTYPE
    // ChaCha20Poly1305
    CHACHA20POLY1305_VPNENCRYPTIONALGORITHMTYPE
)

func (i VpnEncryptionAlgorithmType) String() string {
    return []string{"aes256", "des", "tripleDes", "aes128", "aes128Gcm", "aes256Gcm", "aes192", "aes192Gcm", "chaCha20Poly1305"}[i]
}
func ParseVpnEncryptionAlgorithmType(v string) (interface{}, error) {
    result := AES256_VPNENCRYPTIONALGORITHMTYPE
    switch v {
        case "aes256":
            result = AES256_VPNENCRYPTIONALGORITHMTYPE
        case "des":
            result = DES_VPNENCRYPTIONALGORITHMTYPE
        case "tripleDes":
            result = TRIPLEDES_VPNENCRYPTIONALGORITHMTYPE
        case "aes128":
            result = AES128_VPNENCRYPTIONALGORITHMTYPE
        case "aes128Gcm":
            result = AES128GCM_VPNENCRYPTIONALGORITHMTYPE
        case "aes256Gcm":
            result = AES256GCM_VPNENCRYPTIONALGORITHMTYPE
        case "aes192":
            result = AES192_VPNENCRYPTIONALGORITHMTYPE
        case "aes192Gcm":
            result = AES192GCM_VPNENCRYPTIONALGORITHMTYPE
        case "chaCha20Poly1305":
            result = CHACHA20POLY1305_VPNENCRYPTIONALGORITHMTYPE
        default:
            return 0, errors.New("Unknown VpnEncryptionAlgorithmType value: " + v)
    }
    return &result, nil
}
func SerializeVpnEncryptionAlgorithmType(values []VpnEncryptionAlgorithmType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

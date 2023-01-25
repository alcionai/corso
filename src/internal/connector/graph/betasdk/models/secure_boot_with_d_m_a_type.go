package models
import (
    "errors"
)
// Provides operations to call the add method.
type SecureBootWithDMAType int

const (
    // Not configured, no operation
    NOTCONFIGURED_SECUREBOOTWITHDMATYPE SecureBootWithDMAType = iota
    // Turns on VBS with Secure Boot
    WITHOUTDMA_SECUREBOOTWITHDMATYPE
    // Turns on VBS with Secure Boot and DMA
    WITHDMA_SECUREBOOTWITHDMATYPE
)

func (i SecureBootWithDMAType) String() string {
    return []string{"notConfigured", "withoutDMA", "withDMA"}[i]
}
func ParseSecureBootWithDMAType(v string) (interface{}, error) {
    result := NOTCONFIGURED_SECUREBOOTWITHDMATYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_SECUREBOOTWITHDMATYPE
        case "withoutDMA":
            result = WITHOUTDMA_SECUREBOOTWITHDMATYPE
        case "withDMA":
            result = WITHDMA_SECUREBOOTWITHDMATYPE
        default:
            return 0, errors.New("Unknown SecureBootWithDMAType value: " + v)
    }
    return &result, nil
}
func SerializeSecureBootWithDMAType(values []SecureBootWithDMAType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

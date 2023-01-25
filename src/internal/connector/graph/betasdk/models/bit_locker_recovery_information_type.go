package models
import (
    "errors"
)
// Provides operations to call the add method.
type BitLockerRecoveryInformationType int

const (
    // Store recovery passwords and key packages.
    PASSWORDANDKEY_BITLOCKERRECOVERYINFORMATIONTYPE BitLockerRecoveryInformationType = iota
    // Store recovery passwords only.
    PASSWORDONLY_BITLOCKERRECOVERYINFORMATIONTYPE
)

func (i BitLockerRecoveryInformationType) String() string {
    return []string{"passwordAndKey", "passwordOnly"}[i]
}
func ParseBitLockerRecoveryInformationType(v string) (interface{}, error) {
    result := PASSWORDANDKEY_BITLOCKERRECOVERYINFORMATIONTYPE
    switch v {
        case "passwordAndKey":
            result = PASSWORDANDKEY_BITLOCKERRECOVERYINFORMATIONTYPE
        case "passwordOnly":
            result = PASSWORDONLY_BITLOCKERRECOVERYINFORMATIONTYPE
        default:
            return 0, errors.New("Unknown BitLockerRecoveryInformationType value: " + v)
    }
    return &result, nil
}
func SerializeBitLockerRecoveryInformationType(values []BitLockerRecoveryInformationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

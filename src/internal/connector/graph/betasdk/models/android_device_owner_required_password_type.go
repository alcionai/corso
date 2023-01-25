package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidDeviceOwnerRequiredPasswordType int

const (
    // Device default value, no intent.
    DEVICEDEFAULT_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE AndroidDeviceOwnerRequiredPasswordType = iota
    // There must be a password set, but there are no restrictions on type.
    REQUIRED_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
    // At least numeric.
    NUMERIC_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
    // At least numeric with no repeating or ordered sequences.
    NUMERICCOMPLEX_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
    // At least alphabetic password.
    ALPHABETIC_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
    // At least alphanumeric password
    ALPHANUMERIC_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
    // At least alphanumeric with symbols.
    ALPHANUMERICWITHSYMBOLS_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
    // Low security biometrics based password required.
    LOWSECURITYBIOMETRIC_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
    // Custom password set by the admin.
    CUSTOMPASSWORD_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
)

func (i AndroidDeviceOwnerRequiredPasswordType) String() string {
    return []string{"deviceDefault", "required", "numeric", "numericComplex", "alphabetic", "alphanumeric", "alphanumericWithSymbols", "lowSecurityBiometric", "customPassword"}[i]
}
func ParseAndroidDeviceOwnerRequiredPasswordType(v string) (interface{}, error) {
    result := DEVICEDEFAULT_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
        case "required":
            result = REQUIRED_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
        case "numeric":
            result = NUMERIC_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
        case "numericComplex":
            result = NUMERICCOMPLEX_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
        case "alphabetic":
            result = ALPHABETIC_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
        case "alphanumeric":
            result = ALPHANUMERIC_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
        case "alphanumericWithSymbols":
            result = ALPHANUMERICWITHSYMBOLS_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
        case "lowSecurityBiometric":
            result = LOWSECURITYBIOMETRIC_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
        case "customPassword":
            result = CUSTOMPASSWORD_ANDROIDDEVICEOWNERREQUIREDPASSWORDTYPE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerRequiredPasswordType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerRequiredPasswordType(values []AndroidDeviceOwnerRequiredPasswordType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

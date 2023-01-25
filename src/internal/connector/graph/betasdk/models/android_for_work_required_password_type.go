package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidForWorkRequiredPasswordType int

const (
    // Device default value, no intent.
    DEVICEDEFAULT_ANDROIDFORWORKREQUIREDPASSWORDTYPE AndroidForWorkRequiredPasswordType = iota
    // Low security biometrics based password required.
    LOWSECURITYBIOMETRIC_ANDROIDFORWORKREQUIREDPASSWORDTYPE
    // Required.
    REQUIRED_ANDROIDFORWORKREQUIREDPASSWORDTYPE
    // At least numeric password required.
    ATLEASTNUMERIC_ANDROIDFORWORKREQUIREDPASSWORDTYPE
    // Numeric complex password required.
    NUMERICCOMPLEX_ANDROIDFORWORKREQUIREDPASSWORDTYPE
    // At least alphabetic password required.
    ATLEASTALPHABETIC_ANDROIDFORWORKREQUIREDPASSWORDTYPE
    // At least alphanumeric password required.
    ATLEASTALPHANUMERIC_ANDROIDFORWORKREQUIREDPASSWORDTYPE
    // At least alphanumeric with symbols password required.
    ALPHANUMERICWITHSYMBOLS_ANDROIDFORWORKREQUIREDPASSWORDTYPE
)

func (i AndroidForWorkRequiredPasswordType) String() string {
    return []string{"deviceDefault", "lowSecurityBiometric", "required", "atLeastNumeric", "numericComplex", "atLeastAlphabetic", "atLeastAlphanumeric", "alphanumericWithSymbols"}[i]
}
func ParseAndroidForWorkRequiredPasswordType(v string) (interface{}, error) {
    result := DEVICEDEFAULT_ANDROIDFORWORKREQUIREDPASSWORDTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_ANDROIDFORWORKREQUIREDPASSWORDTYPE
        case "lowSecurityBiometric":
            result = LOWSECURITYBIOMETRIC_ANDROIDFORWORKREQUIREDPASSWORDTYPE
        case "required":
            result = REQUIRED_ANDROIDFORWORKREQUIREDPASSWORDTYPE
        case "atLeastNumeric":
            result = ATLEASTNUMERIC_ANDROIDFORWORKREQUIREDPASSWORDTYPE
        case "numericComplex":
            result = NUMERICCOMPLEX_ANDROIDFORWORKREQUIREDPASSWORDTYPE
        case "atLeastAlphabetic":
            result = ATLEASTALPHABETIC_ANDROIDFORWORKREQUIREDPASSWORDTYPE
        case "atLeastAlphanumeric":
            result = ATLEASTALPHANUMERIC_ANDROIDFORWORKREQUIREDPASSWORDTYPE
        case "alphanumericWithSymbols":
            result = ALPHANUMERICWITHSYMBOLS_ANDROIDFORWORKREQUIREDPASSWORDTYPE
        default:
            return 0, errors.New("Unknown AndroidForWorkRequiredPasswordType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkRequiredPasswordType(values []AndroidForWorkRequiredPasswordType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

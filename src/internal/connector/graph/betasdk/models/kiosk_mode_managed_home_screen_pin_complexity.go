package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type KioskModeManagedHomeScreenPinComplexity int

const (
    // Not configured.
    NOTCONFIGURED_KIOSKMODEMANAGEDHOMESCREENPINCOMPLEXITY KioskModeManagedHomeScreenPinComplexity = iota
    // Numeric values only.
    SIMPLE_KIOSKMODEMANAGEDHOMESCREENPINCOMPLEXITY
    // Alphanumerical value.
    COMPLEX_KIOSKMODEMANAGEDHOMESCREENPINCOMPLEXITY
)

func (i KioskModeManagedHomeScreenPinComplexity) String() string {
    return []string{"notConfigured", "simple", "complex"}[i]
}
func ParseKioskModeManagedHomeScreenPinComplexity(v string) (interface{}, error) {
    result := NOTCONFIGURED_KIOSKMODEMANAGEDHOMESCREENPINCOMPLEXITY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_KIOSKMODEMANAGEDHOMESCREENPINCOMPLEXITY
        case "simple":
            result = SIMPLE_KIOSKMODEMANAGEDHOMESCREENPINCOMPLEXITY
        case "complex":
            result = COMPLEX_KIOSKMODEMANAGEDHOMESCREENPINCOMPLEXITY
        default:
            return 0, errors.New("Unknown KioskModeManagedHomeScreenPinComplexity value: " + v)
    }
    return &result, nil
}
func SerializeKioskModeManagedHomeScreenPinComplexity(values []KioskModeManagedHomeScreenPinComplexity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

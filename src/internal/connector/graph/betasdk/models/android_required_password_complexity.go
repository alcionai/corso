package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidRequiredPasswordComplexity int

const (
    // Device default value, no password.
    NONE_ANDROIDREQUIREDPASSWORDCOMPLEXITY AndroidRequiredPasswordComplexity = iota
    // The required password complexity on the device is of type low as defined by the Android documentation.
    LOW_ANDROIDREQUIREDPASSWORDCOMPLEXITY
    // The required password complexity on the device is of type medium as defined by the Android documentation.
    MEDIUM_ANDROIDREQUIREDPASSWORDCOMPLEXITY
    // The required password complexity on the device is of type high as defined by the Android documentation.
    HIGH_ANDROIDREQUIREDPASSWORDCOMPLEXITY
)

func (i AndroidRequiredPasswordComplexity) String() string {
    return []string{"none", "low", "medium", "high"}[i]
}
func ParseAndroidRequiredPasswordComplexity(v string) (interface{}, error) {
    result := NONE_ANDROIDREQUIREDPASSWORDCOMPLEXITY
    switch v {
        case "none":
            result = NONE_ANDROIDREQUIREDPASSWORDCOMPLEXITY
        case "low":
            result = LOW_ANDROIDREQUIREDPASSWORDCOMPLEXITY
        case "medium":
            result = MEDIUM_ANDROIDREQUIREDPASSWORDCOMPLEXITY
        case "high":
            result = HIGH_ANDROIDREQUIREDPASSWORDCOMPLEXITY
        default:
            return 0, errors.New("Unknown AndroidRequiredPasswordComplexity value: " + v)
    }
    return &result, nil
}
func SerializeAndroidRequiredPasswordComplexity(values []AndroidRequiredPasswordComplexity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

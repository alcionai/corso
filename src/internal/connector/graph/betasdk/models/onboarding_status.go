package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type OnboardingStatus int

const (
    // Unknown
    UNKNOWN_ONBOARDINGSTATUS OnboardingStatus = iota
    // In progress
    INPROGRESS_ONBOARDINGSTATUS
    // Onboarded
    ONBOARDED_ONBOARDINGSTATUS
    // Failed
    FAILED_ONBOARDINGSTATUS
    // Offboarding
    OFFBOARDING_ONBOARDINGSTATUS
    // UnknownFutureValue
    UNKNOWNFUTUREVALUE_ONBOARDINGSTATUS
)

func (i OnboardingStatus) String() string {
    return []string{"unknown", "inprogress", "onboarded", "failed", "offboarding", "unknownFutureValue"}[i]
}
func ParseOnboardingStatus(v string) (interface{}, error) {
    result := UNKNOWN_ONBOARDINGSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_ONBOARDINGSTATUS
        case "inprogress":
            result = INPROGRESS_ONBOARDINGSTATUS
        case "onboarded":
            result = ONBOARDED_ONBOARDINGSTATUS
        case "failed":
            result = FAILED_ONBOARDINGSTATUS
        case "offboarding":
            result = OFFBOARDING_ONBOARDINGSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ONBOARDINGSTATUS
        default:
            return 0, errors.New("Unknown OnboardingStatus value: " + v)
    }
    return &result, nil
}
func SerializeOnboardingStatus(values []OnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

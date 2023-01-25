package models
import (
    "errors"
)
// Provides operations to call the add method.
type DefenderSubmitSamplesConsentType int

const (
    // Send safe samples automatically
    SENDSAFESAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE DefenderSubmitSamplesConsentType = iota
    // Always prompt
    ALWAYSPROMPT_DEFENDERSUBMITSAMPLESCONSENTTYPE
    // Never send
    NEVERSEND_DEFENDERSUBMITSAMPLESCONSENTTYPE
    // Send all samples automatically
    SENDALLSAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE
)

func (i DefenderSubmitSamplesConsentType) String() string {
    return []string{"sendSafeSamplesAutomatically", "alwaysPrompt", "neverSend", "sendAllSamplesAutomatically"}[i]
}
func ParseDefenderSubmitSamplesConsentType(v string) (interface{}, error) {
    result := SENDSAFESAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE
    switch v {
        case "sendSafeSamplesAutomatically":
            result = SENDSAFESAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE
        case "alwaysPrompt":
            result = ALWAYSPROMPT_DEFENDERSUBMITSAMPLESCONSENTTYPE
        case "neverSend":
            result = NEVERSEND_DEFENDERSUBMITSAMPLESCONSENTTYPE
        case "sendAllSamplesAutomatically":
            result = SENDALLSAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE
        default:
            return 0, errors.New("Unknown DefenderSubmitSamplesConsentType value: " + v)
    }
    return &result, nil
}
func SerializeDefenderSubmitSamplesConsentType(values []DefenderSubmitSamplesConsentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

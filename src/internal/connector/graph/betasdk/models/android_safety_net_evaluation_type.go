package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidSafetyNetEvaluationType int

const (
    // Default value. Typical measurements and reference data were used.
    BASIC_ANDROIDSAFETYNETEVALUATIONTYPE AndroidSafetyNetEvaluationType = iota
    // Hardware-backed security features (such as Key Attestation) were used.
    HARDWAREBACKED_ANDROIDSAFETYNETEVALUATIONTYPE
)

func (i AndroidSafetyNetEvaluationType) String() string {
    return []string{"basic", "hardwareBacked"}[i]
}
func ParseAndroidSafetyNetEvaluationType(v string) (interface{}, error) {
    result := BASIC_ANDROIDSAFETYNETEVALUATIONTYPE
    switch v {
        case "basic":
            result = BASIC_ANDROIDSAFETYNETEVALUATIONTYPE
        case "hardwareBacked":
            result = HARDWAREBACKED_ANDROIDSAFETYNETEVALUATIONTYPE
        default:
            return 0, errors.New("Unknown AndroidSafetyNetEvaluationType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidSafetyNetEvaluationType(values []AndroidSafetyNetEvaluationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

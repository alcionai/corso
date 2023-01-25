package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type EligibilityFilteringEnabledEntities int

const (
    NONE_ELIGIBILITYFILTERINGENABLEDENTITIES EligibilityFilteringEnabledEntities = iota
    SWAPREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
    OFFERSHIFTREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
    UNKNOWNFUTUREVALUE_ELIGIBILITYFILTERINGENABLEDENTITIES
    TIMEOFFREASON_ELIGIBILITYFILTERINGENABLEDENTITIES
)

func (i EligibilityFilteringEnabledEntities) String() string {
    return []string{"none", "swapRequest", "offerShiftRequest", "unknownFutureValue", "timeOffReason"}[i]
}
func ParseEligibilityFilteringEnabledEntities(v string) (interface{}, error) {
    result := NONE_ELIGIBILITYFILTERINGENABLEDENTITIES
    switch v {
        case "none":
            result = NONE_ELIGIBILITYFILTERINGENABLEDENTITIES
        case "swapRequest":
            result = SWAPREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
        case "offerShiftRequest":
            result = OFFERSHIFTREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ELIGIBILITYFILTERINGENABLEDENTITIES
        case "timeOffReason":
            result = TIMEOFFREASON_ELIGIBILITYFILTERINGENABLEDENTITIES
        default:
            return 0, errors.New("Unknown EligibilityFilteringEnabledEntities value: " + v)
    }
    return &result, nil
}
func SerializeEligibilityFilteringEnabledEntities(values []EligibilityFilteringEnabledEntities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ComanagementEligibleType int

const (
    COMANAGED_COMANAGEMENTELIGIBLETYPE ComanagementEligibleType = iota
    ELIGIBLE_COMANAGEMENTELIGIBLETYPE
    ELIGIBLEBUTNOTAZUREADJOINED_COMANAGEMENTELIGIBLETYPE
    NEEDSOSUPDATE_COMANAGEMENTELIGIBLETYPE
    INELIGIBLE_COMANAGEMENTELIGIBLETYPE
    // Devices scheduled for Co-Management enrollment
    SCHEDULEDFORENROLLMENT_COMANAGEMENTELIGIBLETYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_COMANAGEMENTELIGIBLETYPE
)

func (i ComanagementEligibleType) String() string {
    return []string{"comanaged", "eligible", "eligibleButNotAzureAdJoined", "needsOsUpdate", "ineligible", "scheduledForEnrollment", "unknownFutureValue"}[i]
}
func ParseComanagementEligibleType(v string) (interface{}, error) {
    result := COMANAGED_COMANAGEMENTELIGIBLETYPE
    switch v {
        case "comanaged":
            result = COMANAGED_COMANAGEMENTELIGIBLETYPE
        case "eligible":
            result = ELIGIBLE_COMANAGEMENTELIGIBLETYPE
        case "eligibleButNotAzureAdJoined":
            result = ELIGIBLEBUTNOTAZUREADJOINED_COMANAGEMENTELIGIBLETYPE
        case "needsOsUpdate":
            result = NEEDSOSUPDATE_COMANAGEMENTELIGIBLETYPE
        case "ineligible":
            result = INELIGIBLE_COMANAGEMENTELIGIBLETYPE
        case "scheduledForEnrollment":
            result = SCHEDULEDFORENROLLMENT_COMANAGEMENTELIGIBLETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_COMANAGEMENTELIGIBLETYPE
        default:
            return 0, errors.New("Unknown ComanagementEligibleType value: " + v)
    }
    return &result, nil
}
func SerializeComanagementEligibleType(values []ComanagementEligibleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

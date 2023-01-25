package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type MembershipRuleProcessingStatusDetails int

const (
    NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS MembershipRuleProcessingStatusDetails = iota
    RUNNING_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    FAILED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    SUCCEEDED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    UNSUPPORTEDFUTUREVALUE_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
)

func (i MembershipRuleProcessingStatusDetails) String() string {
    return []string{"NotStarted", "Running", "Failed", "Succeeded", "UnsupportedFutureValue"}[i]
}
func ParseMembershipRuleProcessingStatusDetails(v string) (interface{}, error) {
    result := NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    switch v {
        case "NotStarted":
            result = NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "Running":
            result = RUNNING_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "Failed":
            result = FAILED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "Succeeded":
            result = SUCCEEDED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "UnsupportedFutureValue":
            result = UNSUPPORTEDFUTUREVALUE_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        default:
            return 0, errors.New("Unknown MembershipRuleProcessingStatusDetails value: " + v)
    }
    return &result, nil
}
func SerializeMembershipRuleProcessingStatusDetails(values []MembershipRuleProcessingStatusDetails) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

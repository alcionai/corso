package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DecisionItemPrincipalResourceMembershipType int

const (
    DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE DecisionItemPrincipalResourceMembershipType = iota
    INDIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
    UNKNOWNFUTUREVALUE_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
)

func (i DecisionItemPrincipalResourceMembershipType) String() string {
    return []string{"direct", "indirect", "unknownFutureValue"}[i]
}
func ParseDecisionItemPrincipalResourceMembershipType(v string) (interface{}, error) {
    result := DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
    switch v {
        case "direct":
            result = DIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
        case "indirect":
            result = INDIRECT_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DECISIONITEMPRINCIPALRESOURCEMEMBERSHIPTYPE
        default:
            return 0, errors.New("Unknown DecisionItemPrincipalResourceMembershipType value: " + v)
    }
    return &result, nil
}
func SerializeDecisionItemPrincipalResourceMembershipType(values []DecisionItemPrincipalResourceMembershipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

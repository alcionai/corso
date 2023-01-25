package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PersonRelationship int

const (
    MANAGER_PERSONRELATIONSHIP PersonRelationship = iota
    COLLEAGUE_PERSONRELATIONSHIP
    DIRECTREPORT_PERSONRELATIONSHIP
    DOTLINEREPORT_PERSONRELATIONSHIP
    ASSISTANT_PERSONRELATIONSHIP
    DOTLINEMANAGER_PERSONRELATIONSHIP
    ALTERNATECONTACT_PERSONRELATIONSHIP
    FRIEND_PERSONRELATIONSHIP
    SPOUSE_PERSONRELATIONSHIP
    SIBLING_PERSONRELATIONSHIP
    CHILD_PERSONRELATIONSHIP
    PARENT_PERSONRELATIONSHIP
    SPONSOR_PERSONRELATIONSHIP
    EMERGENCYCONTACT_PERSONRELATIONSHIP
    OTHER_PERSONRELATIONSHIP
    UNKNOWNFUTUREVALUE_PERSONRELATIONSHIP
)

func (i PersonRelationship) String() string {
    return []string{"manager", "colleague", "directReport", "dotLineReport", "assistant", "dotLineManager", "alternateContact", "friend", "spouse", "sibling", "child", "parent", "sponsor", "emergencyContact", "other", "unknownFutureValue"}[i]
}
func ParsePersonRelationship(v string) (interface{}, error) {
    result := MANAGER_PERSONRELATIONSHIP
    switch v {
        case "manager":
            result = MANAGER_PERSONRELATIONSHIP
        case "colleague":
            result = COLLEAGUE_PERSONRELATIONSHIP
        case "directReport":
            result = DIRECTREPORT_PERSONRELATIONSHIP
        case "dotLineReport":
            result = DOTLINEREPORT_PERSONRELATIONSHIP
        case "assistant":
            result = ASSISTANT_PERSONRELATIONSHIP
        case "dotLineManager":
            result = DOTLINEMANAGER_PERSONRELATIONSHIP
        case "alternateContact":
            result = ALTERNATECONTACT_PERSONRELATIONSHIP
        case "friend":
            result = FRIEND_PERSONRELATIONSHIP
        case "spouse":
            result = SPOUSE_PERSONRELATIONSHIP
        case "sibling":
            result = SIBLING_PERSONRELATIONSHIP
        case "child":
            result = CHILD_PERSONRELATIONSHIP
        case "parent":
            result = PARENT_PERSONRELATIONSHIP
        case "sponsor":
            result = SPONSOR_PERSONRELATIONSHIP
        case "emergencyContact":
            result = EMERGENCYCONTACT_PERSONRELATIONSHIP
        case "other":
            result = OTHER_PERSONRELATIONSHIP
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PERSONRELATIONSHIP
        default:
            return 0, errors.New("Unknown PersonRelationship value: " + v)
    }
    return &result, nil
}
func SerializePersonRelationship(values []PersonRelationship) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

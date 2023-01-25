package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type JoinType int

const (
    // Unknown enrollment join type.
    UNKNOWN_JOINTYPE JoinType = iota
    // The device is joined by Azure AD.
    AZUREADJOINED_JOINTYPE
    // The device is registered by Azure AD.
    AZUREADREGISTERED_JOINTYPE
    // The device is joined by hybrid Azure AD.
    HYBRIDAZUREADJOINED_JOINTYPE
)

func (i JoinType) String() string {
    return []string{"unknown", "azureADJoined", "azureADRegistered", "hybridAzureADJoined"}[i]
}
func ParseJoinType(v string) (interface{}, error) {
    result := UNKNOWN_JOINTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_JOINTYPE
        case "azureADJoined":
            result = AZUREADJOINED_JOINTYPE
        case "azureADRegistered":
            result = AZUREADREGISTERED_JOINTYPE
        case "hybridAzureADJoined":
            result = HYBRIDAZUREADJOINED_JOINTYPE
        default:
            return 0, errors.New("Unknown JoinType value: " + v)
    }
    return &result, nil
}
func SerializeJoinType(values []JoinType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

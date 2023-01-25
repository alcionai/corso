package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type NetworkSingleSignOnType int

const (
    // Disabled
    DISABLED_NETWORKSINGLESIGNONTYPE NetworkSingleSignOnType = iota
    // Pre-Logon
    PRELOGON_NETWORKSINGLESIGNONTYPE
    // Post-Logon
    POSTLOGON_NETWORKSINGLESIGNONTYPE
)

func (i NetworkSingleSignOnType) String() string {
    return []string{"disabled", "prelogon", "postlogon"}[i]
}
func ParseNetworkSingleSignOnType(v string) (interface{}, error) {
    result := DISABLED_NETWORKSINGLESIGNONTYPE
    switch v {
        case "disabled":
            result = DISABLED_NETWORKSINGLESIGNONTYPE
        case "prelogon":
            result = PRELOGON_NETWORKSINGLESIGNONTYPE
        case "postlogon":
            result = POSTLOGON_NETWORKSINGLESIGNONTYPE
        default:
            return 0, errors.New("Unknown NetworkSingleSignOnType value: " + v)
    }
    return &result, nil
}
func SerializeNetworkSingleSignOnType(values []NetworkSingleSignOnType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

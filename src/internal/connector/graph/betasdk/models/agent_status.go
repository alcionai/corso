package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AgentStatus int

const (
    ACTIVE_AGENTSTATUS AgentStatus = iota
    INACTIVE_AGENTSTATUS
)

func (i AgentStatus) String() string {
    return []string{"active", "inactive"}[i]
}
func ParseAgentStatus(v string) (interface{}, error) {
    result := ACTIVE_AGENTSTATUS
    switch v {
        case "active":
            result = ACTIVE_AGENTSTATUS
        case "inactive":
            result = INACTIVE_AGENTSTATUS
        default:
            return 0, errors.New("Unknown AgentStatus value: " + v)
    }
    return &result, nil
}
func SerializeAgentStatus(values []AgentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

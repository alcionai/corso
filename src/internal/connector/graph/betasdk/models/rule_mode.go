package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RuleMode int

const (
    AUDIT_RULEMODE RuleMode = iota
    AUDITANDNOTIFY_RULEMODE
    ENFORCE_RULEMODE
    PENDINGDELETION_RULEMODE
    TEST_RULEMODE
)

func (i RuleMode) String() string {
    return []string{"audit", "auditAndNotify", "enforce", "pendingDeletion", "test"}[i]
}
func ParseRuleMode(v string) (interface{}, error) {
    result := AUDIT_RULEMODE
    switch v {
        case "audit":
            result = AUDIT_RULEMODE
        case "auditAndNotify":
            result = AUDITANDNOTIFY_RULEMODE
        case "enforce":
            result = ENFORCE_RULEMODE
        case "pendingDeletion":
            result = PENDINGDELETION_RULEMODE
        case "test":
            result = TEST_RULEMODE
        default:
            return 0, errors.New("Unknown RuleMode value: " + v)
    }
    return &result, nil
}
func SerializeRuleMode(values []RuleMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

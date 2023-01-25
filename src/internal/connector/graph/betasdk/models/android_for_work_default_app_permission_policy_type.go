package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidForWorkDefaultAppPermissionPolicyType int

const (
    // Device default value, no intent.
    DEVICEDEFAULT_ANDROIDFORWORKDEFAULTAPPPERMISSIONPOLICYTYPE AndroidForWorkDefaultAppPermissionPolicyType = iota
    // Prompt.
    PROMPT_ANDROIDFORWORKDEFAULTAPPPERMISSIONPOLICYTYPE
    // Auto grant.
    AUTOGRANT_ANDROIDFORWORKDEFAULTAPPPERMISSIONPOLICYTYPE
    // Auto deny.
    AUTODENY_ANDROIDFORWORKDEFAULTAPPPERMISSIONPOLICYTYPE
)

func (i AndroidForWorkDefaultAppPermissionPolicyType) String() string {
    return []string{"deviceDefault", "prompt", "autoGrant", "autoDeny"}[i]
}
func ParseAndroidForWorkDefaultAppPermissionPolicyType(v string) (interface{}, error) {
    result := DEVICEDEFAULT_ANDROIDFORWORKDEFAULTAPPPERMISSIONPOLICYTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_ANDROIDFORWORKDEFAULTAPPPERMISSIONPOLICYTYPE
        case "prompt":
            result = PROMPT_ANDROIDFORWORKDEFAULTAPPPERMISSIONPOLICYTYPE
        case "autoGrant":
            result = AUTOGRANT_ANDROIDFORWORKDEFAULTAPPPERMISSIONPOLICYTYPE
        case "autoDeny":
            result = AUTODENY_ANDROIDFORWORKDEFAULTAPPPERMISSIONPOLICYTYPE
        default:
            return 0, errors.New("Unknown AndroidForWorkDefaultAppPermissionPolicyType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkDefaultAppPermissionPolicyType(values []AndroidForWorkDefaultAppPermissionPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

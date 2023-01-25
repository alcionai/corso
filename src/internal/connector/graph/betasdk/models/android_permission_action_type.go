package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidPermissionActionType int

const (
    PROMPT_ANDROIDPERMISSIONACTIONTYPE AndroidPermissionActionType = iota
    AUTOGRANT_ANDROIDPERMISSIONACTIONTYPE
    AUTODENY_ANDROIDPERMISSIONACTIONTYPE
)

func (i AndroidPermissionActionType) String() string {
    return []string{"prompt", "autoGrant", "autoDeny"}[i]
}
func ParseAndroidPermissionActionType(v string) (interface{}, error) {
    result := PROMPT_ANDROIDPERMISSIONACTIONTYPE
    switch v {
        case "prompt":
            result = PROMPT_ANDROIDPERMISSIONACTIONTYPE
        case "autoGrant":
            result = AUTOGRANT_ANDROIDPERMISSIONACTIONTYPE
        case "autoDeny":
            result = AUTODENY_ANDROIDPERMISSIONACTIONTYPE
        default:
            return 0, errors.New("Unknown AndroidPermissionActionType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidPermissionActionType(values []AndroidPermissionActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
